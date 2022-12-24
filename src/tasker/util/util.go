package util

import (
	"football/lib/scheduler"
	"football/tasker/config"
)

var Master *scheduler.Master

func InitMaster() {
	Master = scheduler.NewMaster(config.GetConfig().IProviders)
	Master.Start()
}
