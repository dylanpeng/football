package team

import (
	"fmt"
	"football/lib/scheduler"
	"time"
)

type TeamProvider struct {
	*scheduler.Provider
}

func (m *TeamProvider) Run() {
	fmt.Println("team schedule:", time.Now())

	return
}
