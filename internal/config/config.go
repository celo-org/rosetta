package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/viper"
)

const (
	DatadirEnvVariable = "ROSETTA_DATADIR"
	FornoRC0Url        = "https://rc0-forno.celo-testnet.org/"
	FornoAlfajoresUrl  = "https://alfajores-forno.celo-testnet.org/"
)

type HttpServerConfig struct {
	RequestTimeout time.Duration
	Port           int
	Interface      string
}

func (hs *HttpServerConfig) ListenAddress() string {
	return fmt.Sprintf("%s:%d", hs.Interface, hs.Port)
}

type NodeConfig struct {
	Uri string
}

type ChainConfig struct {
	ChainId   *big.Int
	EpochSize uint64
}

var HttpServer HttpServerConfig
var Node NodeConfig
var Chain ChainConfig

var DataDir string

func ReadConfig() {
	configureDefaults()

	viper.SetEnvPrefix("ROSETTA")
	viper.BindEnv("datadir")

	viper.SetConfigName("rosetta-cfg")
	viper.AddConfigPath(DataDir)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn("No Config File, using defaults")
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

	}

	// httpServer
	HttpServer.RequestTimeout, err = time.ParseDuration(viper.GetString("httpServer.requestTimeout"))
	if err != nil {
		panic(fmt.Errorf("Config: Error reading: %s err: %s\n", "httpServer.requestTimeout", err))
	}
	HttpServer.Port = viper.GetInt("httpServer.port")

	// celo node
	Node.Uri = viper.GetString("node.uri")
}

func SetupDatadir() {
	dataDir, ok := os.LookupEnv(DatadirEnvVariable)
	if !ok {
		log.Crit(fmt.Sprintf("Missing Datadir parameter. Set %s env variable", DatadirEnvVariable))
	}
	dataDir, err := filepath.Abs(dataDir)
	if err != nil {
		log.Crit("Can't resolve Datadir path", "datadir", dataDir)
	}

	stat, err := os.Stat(dataDir)
	switch {
	case err != nil:
		log.Crit("Can't access datadir", "datadir", dataDir, "err", err)
	case !stat.IsDir():
		log.Crit("Datadir is not a directory", "datadir", dataDir)
	}
	DataDir = dataDir
	log.Info("DataDir Configured", "datadir", DataDir)

	readGenesisBlock()
}

func readGenesisBlock() {
	genesisPath := filepath.Join(DataDir, "genesis.json")
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

	Chain.ChainId = new(big.Int).SetUint64(genesis.Config.ChainId)
	Chain.EpochSize = genesis.Config.Isntabul.Epoch
}

func configureDefaults() {
	// httpServer
	viper.SetDefault("httpServer.requestTimeout", "25s")
	viper.SetDefault("httpServer.port", "8080")
	viper.SetDefault("httpServer.interface", "")

	// celo node
	// viper.SetDefault("node.uri", FornoAlfajoresUrl)
	viper.SetDefault("node.uri", FornoRC0Url)
}
