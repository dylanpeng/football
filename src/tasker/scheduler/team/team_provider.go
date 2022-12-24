package team

import (
	"football/common"
	"football/lib/scheduler"
	"time"
)

type TeamProvider struct {
	*scheduler.Provider
}

func (m *TeamProvider) Run() {
	common.Logger.Infof("team schedule:", time.Now())

	return
}
