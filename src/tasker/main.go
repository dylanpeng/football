package main

import (
	"flag"
	"fmt"
	"football/common"
	"football/common/config"
	"football/lib/scheduler"
	"football/tasker/scheduler/match"
	"football/tasker/scheduler/team"
	"football/tasker/util"
	"github.com/BurntSushi/toml"
	"os"
	"os/signal"
	"syscall"
)

var (
	configFile = flag.String("c", "config.toml", "config file path")
)

func main() {
	// parse flag
	flag.Parse()

	conf := &config.Config{}

	m, err := toml.DecodeFile(*configFile, conf)

	if err != nil {
		fmt.Printf("toml decode failed. err: %s | m: %s", err, m)
		return
	}

	_ = conf.Init()

	err = common.InitDB()
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
