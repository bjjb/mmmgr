/*
Package config provides tools for configuring mmmgr.
*/
package cfg

import (
	"github.com/spf13/viper"
	"log"
	"path"
)

// A cfg contains configuration variables and methods
type cfg struct {
	*viper.Viper
}

// The Cfg is the underlying cfg, set up in init()
var Cfg *cfg

func init() {
	c := &cfg{viper.New()}
	c.SetConfigName("config")
	c.AddConfigPath(path.Join("$HOME/.mmmgr"))
	c.SetEnvPrefix("mmmgr")
	c.AutomaticEnv()
	if err := c.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	Cfg = c
}

// UnmarshalKey unmarshals the config variable at k into the object i
func UnmarshalKey(k string, i interface{}) error {
	return Cfg.UnmarshalKey(k, i)
}
