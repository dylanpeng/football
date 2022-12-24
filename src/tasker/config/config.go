package config

import (
	oConf "football/common/config"
	"football/lib/scheduler"
	"github.com/BurntSushi/toml"
)

var conf *Config

type Config struct {
	*oConf.Config
	Providers []*scheduler.Provider `toml:"providers" json:"providers"`
}

func Init(file string) error {
	conf = &Config{
		Config: oConf.Default(),
	}

	if _, err := toml.DecodeFile(file, conf); err != nil {
		return err
	}

	if err := conf.Config.Init(); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return conf
}
