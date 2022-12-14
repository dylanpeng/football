package util

import "football/lib/scheduler"

var Master *scheduler.Master

func InitMaster(providers []scheduler.IProvider) {
	Master = scheduler.NewMaster(providers)
	Master.Start()
}
