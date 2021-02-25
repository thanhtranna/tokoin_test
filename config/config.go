package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigSchema struct {
	Data struct {
		Organization string `json:"organization"`
		User         string `json:"user"`
		Ticket       string `json:"ticket"`
	} `json:"data"`
}

var Config ConfigSchema

func init() {
	config := viper.New()

	// Set config file
	config.SetConfigName("config")
	// Set the path to look for the configurations file
	config.AddConfigPath(".")
	config.AddConfigPath("config/")
	config.AddConfigPath("../config/")
	config.AddConfigPath("../")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	err = config.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
