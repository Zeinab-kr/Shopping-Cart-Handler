package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Data   string `gorm:"type:varchar(2048)" validate:"max=2048"`
	State  string `gorm:"type:varchar(10)" validate:"customState"`
}

type CartInput struct {
	Data  string
	State string
}
