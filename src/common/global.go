package common

import (
	"football/common/config"
	"football/lib/gorm"
	oGorm "gorm.io/gorm"
)

var dbPool *gorm.Pool

func InitDB() (err error) {
	confs := config.GetConfig().DB
	dbPool = gorm.NewPool()

	for k, v := range confs {
		if err := dbPool.Add(k, v); err != nil {
			return err
		}
	}

	return nil
}

func GetDB(name string) (*oGorm.DB, error) {
	return dbPool.Get(name)
}
