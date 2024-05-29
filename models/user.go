package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user of the application
type User struct {
	gorm.Model
	Username string `validate:"required" gorm:"unique"`
	Password string `validate:"required,min=8"`
	Carts  []Cart
}

// HashPassword hashes the user's password using bcrypt
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(bytes)
	return err
}

// CheckPassword checks if the provided password matches the user's hashed password
func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
