package geth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/internal/fileutils"
	"github.com/celo-org/rosetta/service"
	"github.com/ethereum/go-ethereum/log"
)

type gethService struct {
	genesisPath string
	gethBinary  string
	paths       GethPaths
	staticNodes []string
	chainParams *celo.ChainParameters

	cmd     *exec.Cmd
	running service.RunningLock
	logger  log.Logger
	errors  chan error
}

func NewGethService(
	gethBinary string,
	datadir string,
	genesisPath string,
	staticNodes []string) *gethService {
	return &gethService{
		genesisPath: genesisPath,
		gethBinary:  gethBinary,
		paths:       GethPaths(datadir),
		staticNodes: staticNodes,
		logger:      log.New("service", "geth"),
	}
}

func (gs *gethService) IpcFilePath() string {
	return gs.paths.IpcFile()
}

func (gs *gethService) ChainParameters() *celo.ChainParameters {
	return gs.chainParams
}

func (gs *gethService) Name() string {
	return "geth"
}

func (gs *gethService) Running() bool {
	return gs.running.Running()
}

func (gs *gethService) Setup() error {
	if err := os.MkdirAll(gs.paths.Datadir(), os.ModePerm); err != nil {
		return fmt.Errorf("Can't create celo-blockchain datadir: %w", err)
	}

	// Read Genesis to get chain parameters
	gs.chainParams = chainParamsFromGenesisFile(gs.genesisPath)

	if err := gs.ensureGethInit(); err != nil {
		return err
	}

	return nil
}

func (gs *gethService) Start(ctx context.Context) error {
	if err := gs.running.EnableOrFail(); err != nil {
		return err
	}
	defer gs.running.Disable()

	if err := gs.Setup(); err != nil {
		return err
	}

	if err := gs.setupStaticNodes(); err != nil {
		return err
	}

	if err := gs.startGeth(); err != nil {
		return err
	}

	// Listen to stop signal, and kill the process
	go func() {
		<-ctx.Done()
		if err := gs.cmd.Process.Signal(os.Interrupt); err != nil {
			// Not much else to do. Failed to send a signal
			panic(fmt.Errorf("Error sending signal: %w ", err))
		}
	}()

	if err := gs.cmd.Wait(); err != nil {
		return err
	}

	return nil

}

func (gs *gethService) gethCmd(args ...string) *exec.Cmd {
	fullArgs := append([]string{"--datadir", gs.paths.Datadir()}, args...)
	return exec.Command(gs.gethBinary, fullArgs...)
}

func (gs *gethService) setupStaticNodes() error {
	var staticNodesRaw []byte
	var err error

	if staticNodesRaw, err = json.Marshal(gs.staticNodes); err != nil {
		return fmt.Errorf("Can't serialize static nodes: %w", err)
	}

	if err = ioutil.WriteFile(gs.paths.StaticNodesFile(), staticNodesRaw, 0644); err != nil {
		return fmt.Errorf("Can't serialize static nodes: %w", err)
	}

	return nil
}

func (gs *gethService) ensureGethInit() error {
	// Check if geth is initialized already
	flagFile := gs.paths.GethInitializedFile()

	if fileutils.FileExists(flagFile) {
		gs.logger.Info("Geth Already initialized... skipping init")
		return nil
	}

	gs.logger.Info("Running geth init")
	out, err := gs.gethCmd("init", gs.genesisPath).CombinedOutput()
	if err != nil {
		gs.logger.Error("Error running geth init", "err", err)
		fmt.Println(string(out))
		return err
	}
	if err = fileutils.TouchFile(flagFile); err != nil {
		gs.logger.Error("Error creating marker file", "err", err)
		return err
	}
	return nil
}

func (gs *gethService) startGeth() error {
	gethArgs := []string{
		"--verbosity", "3",
		"--syncmode", "full",
		"--networkid", gs.chainParams.ChainId.String(),
		"--rpc",
		"--rpcaddr", "127.0.0.1",
		"--rpcapi", "eth,net,web3,debug,admin,personal",
		"--ws",
		"--light.serve", "0",
		"--light.maxpeers", "0",
		"--maxpeers", "1100",
		"--gcmode", "archive",
		"--consoleformat", "term",
		// "--consoleoutput", "split",
	}

	fmt.Println("geth", strings.Join(gethArgs, " "))
	f, err := os.OpenFile(gs.paths.LogFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		gs.logger.Error("Can't open geth logfile", "err", err)
		return err
	}

	// TODO: fixt this
	// defer f.Close()

	cmd := gs.gethCmd(gethArgs...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	// cmd.Stdout = f
	cmd.Stderr = f

	if err = cmd.Start(); err != nil {
		gs.logger.Error("Error starting geth", "err", err)
		return err
	}

	gs.cmd = cmd
	return nil
}

type GethPaths string

func (g GethPaths) GethInitializedFile() string {
	return filepath.Join(g.Datadir(), ".geth-initialized")
}

func (g GethPaths) LogFile() string {
	return filepath.Join(g.Datadir(), "celo.log")
}

func (g GethPaths) IpcFile() string {
	return filepath.Join(g.Datadir(), "geth.ipc")
}

func (g GethPaths) StaticNodesFile() string {
	return filepath.Join(g.Datadir(), "/Celo/static-nodes.json")
}

func (g GethPaths) Datadir() string {
	return string(g)
}

func chainParamsFromGenesisFile(genesisPath string) *celo.ChainParameters {
	data, err := ioutil.ReadFile(genesisPath)
	if err != nil {
		log.Crit("Can't read genesis.json on DataDir", "genesisPath", genesisPath, "err", err)
	}

	// We only map the fields we need
	var genesis struct {
		Config struct {
			ChainId  uint64 `json:"chainId"`
			Isntabul struct {
				Epoch uint64 `json:"epoch"`
			} `json:"istanbul"`
		} `json:"config"`
	}

	if err = json.Unmarshal(data, &genesis); err != nil {
		log.Crit("Can't parse genesis.json on DataDir", "genesisPath", genesisPath, "err", err)
	}

	return &celo.ChainParameters{
		ChainId:   new(big.Int).SetUint64(genesis.Config.ChainId),
		EpochSize: genesis.Config.Isntabul.Epoch,
	}
}
