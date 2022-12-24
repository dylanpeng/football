package config

import (
	oConf "football/common/config"
	"github.com/BurntSushi/toml"
)

var conf *Config

type Config struct {
	*oConf.Config
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
