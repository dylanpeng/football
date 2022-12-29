package entity

import "fmt"

type League struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	ThirdId    int64  `json:"third_id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func (e *League) TableName() string {
	return "data_league"
}

func (e *League) PrimaryPairs() []interface{} {
	return []interface{}{"id", e.Id}
}

func (e *League) PrimarySeted() bool {
	return e.Id > 0
}

func (e *League) String() string {
	return fmt.Sprintf("%+v", *e)
}
