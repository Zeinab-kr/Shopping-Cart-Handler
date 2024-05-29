package models

import (
	"gorm.io/gorm"
)

// Cart represents a shopping cart
type Cart struct {
	gorm.Model
	UserID uint   `gorm:"not null"` // Link to the User model
	Data   string `gorm:"type:varchar(2048)" validate:"max=2048"`
	State  string `gorm:"type:varchar(10)" validate:"customState"`
}

// CartInput is a struct used to represent the input data for a cart
type CartInput struct {
	Data  string
	State string
}
