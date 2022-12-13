package model

import (
	"football/common"
	"football/common/entity"
	"gorm.io/gorm"
)

type dbBase struct {
	readDbName  string
	writeDbName string
}

func createDBModel(readInstance, writeInstance string) *dbBase {
	return &dbBase{
		readDbName:  readInstance,
		writeDbName: writeInstance,
	}
}

func (d *dbBase) getDB() (db *gorm.DB, err error) {
	return common.GetDb(d.writeDbName)
}

func (d *dbBase) getReadDB() (db *gorm.DB, err error) {
	return common.GetDb(d.readDbName)
}

func (d *dbBase) Add(entity entity.IEntity) (err error) {
	db, err := common.GetDb(d.writeDbName)

	if err != nil {
		return
	}

	err = db.Create(entity).Error
	return
}

func (d *dbBase) Update(entity entity.IEntity, params map[string]interface{}) (err error) {
	db, err := common.GetDb(d.writeDbName)

	if err != nil {
		return
	}

	if params == nil {
		err = db.Save(entity).Error
	} else {
		err = db.Model(entity).Updates(params).Error
	}

	return
}

func (d *dbBase) Get(entity entity.IEntity) (err error) {
	db, err := common.GetDb(d.writeDbName)

	if err != nil {
		return
	}

	err = db.First(entity).Error

	return
}
