package config

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/viper"
)

type HttpServerConfig struct {
	RequestTimeout time.Duration
	Port           int
	Interface      string
}

func (hs *HttpServerConfig) ListenAddress() string {
	return fmt.Sprintf("%s:%d", hs.Interface, hs.Port)
}

var HttpServer HttpServerConfig

func ReadConfig() {

	viper.SetDefault("httpServer.requestTimeout", "25s")

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

	HttpServer.RequestTimeout, err = time.ParseDuration(viper.GetString("httpServer.requestTimeout"))
	if err != nil {
		panic(fmt.Errorf("Config: Error reading: %s \n", "httpServer.requestTimeout"))
	}
	HttpServer.Port = viper.GetInt("httpServer.port")
}
