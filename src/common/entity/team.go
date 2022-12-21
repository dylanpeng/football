package entity

import "fmt"

type Team struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	ThirdId    int64  `json:"third_id"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func (e *Team) TableName() string {
	return "data_team"
}

func (e *Team) PrimarySeted() bool {
	return e.Id > 0
}

func (e *Team) String() string {
	return fmt.Sprintf("%+v", *e)
}
