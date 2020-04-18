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
	"fmt"
	"os"

	"github.com/celo-org/rosetta/cmd/cli"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "rosetta",
	Short: "A standard for blockchain interaction",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.SetEnvPrefix("ROSETTA")
	viper.AutomaticEnv() // read in environment variables that match

	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	RootCmd.AddCommand(cli.CliCmd)

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := homedir.Dir()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// fmt.Println("datadir=", viper.Get("datadir"))

	// 	// Search config in home directory with name ".rosetta" (without extension).
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigName(".rosetta")
	// }

	// // If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// }
}

func exitOnMissingConfig(cmd *cobra.Command, configKey string) {
	if !viper.IsSet(configKey) {
		cmd.Println(cmd.UsageString())
		cmd.Printf("Missing required config: %s\n", configKey)
		os.Exit(1)
	}
}
