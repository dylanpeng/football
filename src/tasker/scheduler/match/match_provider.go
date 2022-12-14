package match

import (
	"fmt"
	"football/lib/leisu"
	"football/lib/scheduler"
	"football/tasker/logic/schedule"
	"time"
)

type MatchProvider struct {
	*scheduler.Provider
}

func (m *MatchProvider) Run() {
	fmt.Println("match schedule:", time.Now())

	scheduleObject, err := leisu.QueryMatch()

	if err != nil {
		fmt.Printf("QueryMatch failed. match: %s | err: %s", scheduleObject, err)
		return
	}

	_ = schedule.Match.InitSchedule(scheduleObject)

	return
}
