/*
Copyright © 2020 Celo Org

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"path/filepath"
	"time"

	"github.com/celo-org/rosetta/api"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:              "serve",
	Short:            "Start rosetta server",
	Args:             cobra.NoArgs,
	PersistentPreRun: validateDatadir,
}

var httpPort uint
var httpAddress string
var requestTimeout time.Duration
var datadir string

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().UintVar(&httpPort, "port", 8080, "Listening port for http server")
	serveCmd.PersistentFlags().StringVar(&httpAddress, "address", "", "Listening address for http server")
	serveCmd.PersistentFlags().DurationVar(&requestTimeout, "reqTimeout", 25*time.Second, "Timeout when serving a request")

	serveCmd.PersistentFlags().StringVar(&datadir, "datadir", "", "datadir to use")
	exitOnError(serveCmd.MarkPersistentFlagDirname("datadir"))
	exitOnError(serveCmd.MarkPersistentFlagRequired("datadir"))
}

func getRosettaServerConfig() *api.RosettaServerConfig {
	return &api.RosettaServerConfig{
		Port:           httpPort,
		Interface:      httpAddress,
		RequestTimeout: requestTimeout,
	}
}

func validateDatadir(cmd *cobra.Command, args []string) {
	absDatadir, err := filepath.Abs(datadir)
	if err != nil {
		log.Crit("Can't resolve datadir path", "datadir", absDatadir)
	}

	stat, err := os.Stat(absDatadir)
	switch {
	case err != nil:
		log.Crit("Can't access datadir", "datadir", absDatadir, "err", err)
	case !stat.IsDir():
		log.Crit("Datadir is not a directory", "datadir", absDatadir)
	}
	datadir = absDatadir
	log.Info("DataDir Configured", "datadir", datadir)
}
