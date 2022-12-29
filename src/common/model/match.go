package model

import (
	"football/common/entity"
	"gorm.io/gorm"
	"time"
)

var Match = &matchModel{
	baseModel: createModel(
		"main-slave",
		"main-master",
		"main-slave",
		"main-master",
		"match",
		time.Minute*10,
	),
}

type matchModel struct {
	*baseModel
}

func (m *matchModel) GetMatchByThirdId(thirdId int64) (match *entity.Match, err error) {
	match = &entity.Match{}

	key := m.Cache.GetKey("third_id", thirdId)

	exist, err := m.Cache.Get(key, match)

	if err != nil {
		return nil, err
	}

	if exist {
		return
	}

	db, err := m.DB.getDB(false)

	if err != nil {
		return
	}

	err = db.Where("match_third_id = ?", thirdId).First(match).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	m.Cache.Set(key, match, m.Cache.expire)

	return
}
