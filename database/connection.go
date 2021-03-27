package database

import (
	"fmt"
	"go-jwt/env"
	"go-jwt/models"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() error {
	// Get necessary environment variables needed to connect to Postgres
	err := godotenv.Load()
	checkError(err)
	host, user, password, port, name := env.GetEnv()

	// Database Connection
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
	// Opening Connction to DB
	db, err := gorm.Open(postgres.Open(dbURI))

	db.Create(&models.User{})

	if err != nil {
		return err
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
