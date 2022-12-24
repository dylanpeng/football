package main

import (
	"flag"
	"fmt"
	"football/common"
	"football/lib/scheduler"
	"football/tasker/config"
	"football/tasker/scheduler/match"
	"football/tasker/scheduler/team"
	"football/tasker/util"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	configFile = flag.String("c", "config.toml", "config file path")
)

func main() {
	// parse flag
	flag.Parse()

	// set max cpu core
	runtime.GOMAXPROCS(runtime.NumCPU())

	// parse config file
	if err := config.Init(*configFile); err != nil {
		log.Fatalf("Fatal Error: can't parse config file!!!\n%s", err)
	}

	// init log
	if err := common.InitLogger(); err != nil {
		log.Fatalf("Fatal Error: can't initialize logger!!!\n%s", err)
	}

	if err := common.InitDB(); err != nil {
		log.Fatalf("Fatal Error: can't initialize db clients!!!\n%s", err)
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
