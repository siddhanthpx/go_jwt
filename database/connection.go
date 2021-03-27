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

var DB *gorm.DB

func ConnectDB() error {
	// Get necessary environment variables needed to connect to Postgres
	err := godotenv.Load()
	checkError(err)
	host, user, password, port, name := env.GetEnv()

	// Database Connection
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
	// Opening Connction to DB
	connection, err := gorm.Open(postgres.Open(dbURI))

	DB = connection

	connection.AutoMigrate(&models.User{})

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
