package entity

type League struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	ThirdId    int64  `json:"third_id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}
