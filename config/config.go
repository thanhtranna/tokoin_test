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
	config.SetConfigFile(`config.json`)
	// Set the path to look for the configurations file
	config.AddConfigPath("./config/") // Optionally look for config in the working directory.
	// config.AddConfigPath("../config/") // Look for config needed for tests.
	// config.AddConfigPath("../")        // Look for config needed for tests.

	config.AutomaticEnv()

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	err = config.Unmarshal(&Config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	// fmt.Printf("Current Config: %+v", Config)
}
