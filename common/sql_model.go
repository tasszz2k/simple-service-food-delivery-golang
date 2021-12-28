package common

import "time"

type SQLModel struct {
	Id       int        `json:"id" gorm:"column:id"`
	Status   int        `json:"status" gorm:"column:status"`
	CreateAt *time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt *time.Time `json:"update_at" gorm:"column:update_at"`
}
