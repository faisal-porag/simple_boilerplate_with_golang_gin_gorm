package db

import "golang_boilerplate_with_gin/models"

// CreateBook >>>>>> Add methods for database operations here
func (repo *PgRepository) CreateBook(book *models.Book) error {
	return repo.db.Create(book).Error
}

func (repo *PgRepository) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	if err := repo.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
