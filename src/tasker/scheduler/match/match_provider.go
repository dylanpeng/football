package match

import (
	"football/common"
	"football/lib/scheduler"
	"football/pkg/leisu"
	"football/tasker/logic/schedule"
	"time"
)

type MatchProvider struct {
	*scheduler.Provider
}

func (m *MatchProvider) Run() {
	common.Logger.Infof("match schedule:", time.Now())

	scheduleObject, err := leisu.QueryMatch()

	if err != nil {
		common.Logger.Errorf("QueryMatch failed. match: %s | err: %s", scheduleObject, err)
		return
	}

	_ = schedule.Match.InitSchedule(scheduleObject)

	return
}
