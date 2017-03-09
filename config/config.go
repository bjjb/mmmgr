/*
Package config contains configuration variables for mmmgr, read from a
YAML/JSON/TOML file in ~/.mmmgr/config, and overridden by environment
variables which are prexied with `MMMGR_`. Uses
https://github.com/spf13/viper.
*/
package config

import (
	"github.com/spf13/viper"
	"log"
)

/*
TVDB contains the credentials for the TVDB client
*/
var TVDB map[string]string

func init() {
	// Set up the configuration
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.mmmgr")
	viper.SetEnvPrefix("mmmgr")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	TVDB = viper.GetStringMapString("tvdb")
}
