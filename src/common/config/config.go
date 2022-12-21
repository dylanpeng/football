package config

import "football/lib/gorm"

type Config struct {
	DB map[string]*gorm.Config
}

var conf *Config

func (c *Config) Init() (err error) {
	conf = c
	return
}

func GetConfig() *Config {
	return conf
}
