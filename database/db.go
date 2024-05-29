package database

import (
	"fmt"
	"Cart/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	// Get the database connection details from the config.go file
    dbConfig := GetDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
						dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("Failed to connect to database: %v\n", err)
	}

	log.Println("Database connection established")

	err = DB.AutoMigrate(&models.User{}, &models.Cart{})
	if err != nil {
		log.Panicf("Failed to migrate the database: %v\n", err)
	}
}
