package config

import (
	"football/lib/gorm"
	"football/lib/logger"
)

type Config struct {
	DB  map[string]*gorm.Config `toml:"db" json:"db"`
	Log *logger.Config          `toml:"log" json:"log"`
}

var conf *Config

func (c *Config) Init() (err error) {
	conf = c
	return
}

func Default() *Config {
	return &Config{
		Log: logger.DefaultConfig(),
	}
}

func GetConfig() *Config {
	return conf
}
