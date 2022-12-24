package util

import (
	"football/lib/scheduler"
	"football/tasker/config"
	"football/tasker/scheduler/match"
	"football/tasker/scheduler/team"
)

var allProviders = map[string]scheduler.IProvider{
	"schedule_match": &match.MatchProvider{Provider: &scheduler.Provider{}},
	"schedule_team":  &team.TeamProvider{Provider: &scheduler.Provider{}},
}
var Master *scheduler.Master

func InitMaster() {
	if len(config.GetConfig().Providers) == 0 {
		return
	}

	providers := make([]scheduler.IProvider, 0)

	for _, conf := range config.GetConfig().Providers {
		p, exist := allProviders[conf.Name]

		if exist {
			p.SetName(conf.Name)
			p.SetCronExpression(conf.CronExpression)
			providers = append(providers, p)
		}
	}

	Master = scheduler.NewMaster(providers)
	Master.Start()
}
