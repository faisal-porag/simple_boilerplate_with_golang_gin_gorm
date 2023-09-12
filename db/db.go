package db

import (
	"github.com/rs/zerolog/log"
	"golang_boilerplate_with_gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgRepository struct {
	db *gorm.DB
}

func NewPgRepository(dbUrl string) (*PgRepository, error) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	// REMOVE WHEN TEST DONE
	errMigrate := db.AutoMigrate(&models.Book{})
	if errMigrate != nil {
		log.Error().Err(errMigrate)
	}

	return &PgRepository{
		db: db,
	}, nil

}
