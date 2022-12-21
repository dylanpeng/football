package model

import (
	"football/common/entity"
	"gorm.io/gorm"
)

var Match = &matchModel{
	baseDBModel: createDBModel(
		"main-slave",
		"main-master",
	),
}

type matchModel struct {
	*baseDBModel
}

func (m *matchModel) GetMatchByThirdId(thirdId int64) (match *entity.Match, err error) {
	db, err := m.getDB(false)

	if err != nil {
		return
	}

	match = &entity.Match{}
	err = db.Where("match_third_id = ?", thirdId).First(match).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return
}
