package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"string;not null"`
	Gender string `json:"gender" gorm:"string;not null"`
	Age    string `json:"age" gorm:"uint8;not null"`
}
