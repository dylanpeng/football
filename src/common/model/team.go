package model

import (
	"football/common/entity"
	"gorm.io/gorm"
)

var Team = &teamModel{
	dbBase: createDBModel(
		"slave",
		"master"),
}

type teamModel struct {
	*dbBase
}

func (t *teamModel) GetTeamByThirdId(thirdId int64) (team *entity.Team, err error) {
	db, err := t.getDB()

	if err != nil {
		return
	}

	team = &entity.Team{}
	err = db.Where("third_id = ?", thirdId).First(team).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return
}
