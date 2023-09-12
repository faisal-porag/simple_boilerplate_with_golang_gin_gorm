package db

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"golang_boilerplate_with_gin/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectPostgres(cfg *config.Config) (*gorm.DB, error) {
	var err error
	host := cfg.DatabaseHost
	username := cfg.DatabaseUser
	password := cfg.DatabasePassword
	databaseName := cfg.DatabaseDbName
	port := cfg.DatabasePort

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error().Err(err)
		return nil, err
	} else {
		log.Info().Msg("Successfully connected to the database")
	}

	return Database, nil
}
