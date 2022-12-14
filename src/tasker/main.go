package main

import (
	"fmt"
	"football/common"
	"football/common/config"
	"football/lib/scheduler"
	"football/tasker/scheduler/match"
	"football/tasker/scheduler/team"
	"football/tasker/util"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.Config{}

	conf.DBConfigs = make([]*config.DBConfig, 0, 8)

	conf.DBConfigs = append(conf.DBConfigs, &config.DBConfig{
		Name:         "master",
		UserName:     "dev",
		Password:     "123!@#qweASD",
		SourceUrl:    "127.0.0.1",
		Port:         "3306",
		DataBaseName: "sport",
	})

	conf.DBConfigs = append(conf.DBConfigs, &config.DBConfig{
		Name:         "slave",
		UserName:     "dev",
		Password:     "123!@#qweASD",
		SourceUrl:    "127.0.0.1",
		Port:         "3306",
		DataBaseName: "sport",
	})

	err := common.InitDB(conf.DBConfigs)
	if err != nil {
		fmt.Printf("Init Db failed. err: %s", err)
		return
	}

	// init schedule
	util.InitMaster([]scheduler.IProvider{
		&match.MatchProvider{&scheduler.Provider{
			Name:           "schedule_match",
			CronExpression: "0 */1 * * * *",
		}},
		&team.TeamProvider{&scheduler.Provider{
			Name:           "schedule_team",
			CronExpression: "* * * * * *",
		}},
	})

	// waitting for exit signal
	exit := make(chan os.Signal, 1)
	stopSigs := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
	signal.Notify(exit, stopSigs...)

	// catch exit signal
	sign := <-exit
	fmt.Printf("stop by exit signal '%s'\n", sign)

	return
}
