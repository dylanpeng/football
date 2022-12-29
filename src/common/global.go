package common

import (
	"football/common/config"
	"football/lib/gorm"
	"football/lib/logger"
	"football/lib/redis"
	oRedis "github.com/go-redis/redis/v9"
	oGorm "gorm.io/gorm"
)

var cachePool *redis.Pool
var dbPool *gorm.Pool
var Logger *logger.Logger

func InitLogger() (err error) {
	conf := config.GetConfig().Log
	Logger, err = logger.NewLogger(conf)
	return err
}

func InitDB() (err error) {
	confs := config.GetConfig().DB
	dbPool = gorm.NewPool(Logger)

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

func InitCache() {
	confs := config.GetConfig().Cache
	cachePool = redis.NewPool()

	for k, v := range confs {
		cachePool.Add(k, v)
	}
}

func GetCache(name string) (*oRedis.Client, error) {
	return cachePool.Get(name)
}
