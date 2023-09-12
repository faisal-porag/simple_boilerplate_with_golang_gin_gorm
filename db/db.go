package db

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectPostgres() *gorm.DB {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error().Err(err)
		return nil
	} else {
		log.Info().Msg("Successfully connected to the database")
	}

	return Database
}
