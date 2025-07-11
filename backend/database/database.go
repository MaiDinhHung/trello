package database

import (
	"fmt"
	"log"
	"os"
	"trello/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db

	// Auto migrate model
	db.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.List{},
		&models.Card{},	
		&models.Comment{},
		&models.Label{},
	)
	fmt.Println("Connected to DB & Migrated")

}
