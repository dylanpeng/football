package main

import (
	"fmt"
	"football/common"
	"football/common/config"
	"football/lib/leisu"
	"football/tasker/logic/schedule"
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

	scheduleObject, err := leisu.QueryMatch()

	if err != nil {
		fmt.Printf("QueryMatch failed. match: %s | err: %s", scheduleObject, err)
		return
	}

	_ = schedule.InitSchedule(scheduleObject)

	return
}
