package db

import (
	"golang_boilerplate_with_gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// REMOVE WHEN TEST DONE
	errMigrate := db.AutoMigrate(&models.Book{})
	if errMigrate != nil {
		log.Fatalln(errMigrate)
	}

	return db, nil
}
