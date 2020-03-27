package config

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/viper"
)

const (
	FornoRC0Url       = "https://rc0-forno.celo-testnet.org/"
	FornoAlfajoresUrl = "https://alfajores-forno.celo-testnet.org/"
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

var HttpServer HttpServerConfig
var Node NodeConfig

func configureDefaults() {
	// httpServer
	viper.SetDefault("httpServer.requestTimeout", "25s")
	viper.SetDefault("httpServer.port", "8080")
	viper.SetDefault("httpServer.interface", "")

	// celo node
	// viper.SetDefault("node.uri", FornoAlfajoresUrl)
	viper.SetDefault("node.uri", FornoRC0Url)
}

func ReadConfig() {
	configureDefaults()

	viper.SetConfigName("config")
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
