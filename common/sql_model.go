package common

import "time"

type SQLModel struct{
	Id int `json:"id" gorm:"column:id"`
	Status int `json:"status" gorm:"column:status"`
	
	CreateddAt *time.Time `json:"created_at" gorm:"column:created_at"`
	
	UpdateAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	
}