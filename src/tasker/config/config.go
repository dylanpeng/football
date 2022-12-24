package config

import (
	oConf "football/common/config"
	"football/lib/scheduler"
	"football/tasker/scheduler/match"
	"football/tasker/scheduler/team"
	"github.com/BurntSushi/toml"
	"reflect"
)

var conf *Config

type Config struct {
	*oConf.Config
	Providers  *Providers            `toml:"providers" json:"providers"`
	IProviders []scheduler.IProvider `toml:"-" json:"-"`
}

type Providers struct {
	MatchProvider *match.MatchProvider `toml:"match_provider" json:"match_provider"`
	TeamProvider  *team.TeamProvider   `toml:"team_provider" json:"team_provider"`
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

	if conf.Providers != nil {
		n, v := reflect.TypeOf(*conf.Providers).NumField(), reflect.ValueOf(*conf.Providers)
		conf.IProviders = make([]scheduler.IProvider, 0, n)

		for i := 0; i < n; i++ {
			vv := v.Field(i)

			if vv.IsNil() {
				continue
			}

			p, ok := vv.Interface().(scheduler.IProvider)

			if !ok {
				continue
			}

			conf.IProviders = append(conf.IProviders, p)
		}
	}

	return nil
}

func GetConfig() *Config {
	return conf
}
