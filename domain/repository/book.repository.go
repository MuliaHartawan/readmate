package repository

import (
	"fmt"
	"log"
	"readmate/delivery/http/dto/request"
	"readmate/domain/entity"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *request.RequestBookCreateDTO) error
	GetAll() ([]*entity.Book, error)
	GetByID(id uint) (*entity.Book, error)
	Update(id uint, book *request.RequestBookUpdateDTO) error
	Delete(id uint) error
}

type bookRepository struct {
	database *gorm.DB
}

func NewBookRepository(database *gorm.DB) *bookRepository {
	return &bookRepository{
		database: database,
	}
}

func (r *bookRepository) Create(book *request.RequestBookCreateDTO) error {
	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Genre:  book.Genre,
	}

	result := r.database.Create(&newBook)
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error create book: %v", result.Error))
		return result.Error
	}

	return nil
}

func (r *bookRepository) GetAll() ([]*entity.Book, error) {
	var books []*entity.Book

	result := r.database.Find(&books)
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error get all book : %v", result.Error))
		return nil, result.Error
	}

	return books, nil
}

func (r *bookRepository) GetByID(id uint) (*entity.Book, error) {
	var book *entity.Book

	result := r.database.First(&book)
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error get book by ID: %v", result.Error))
		return nil, result.Error
	}

	return book, nil
}

func (r *bookRepository) Update(id uint, req *request.RequestBookUpdateDTO) error {
	var book *entity.Book

	result := r.database.First(&book, id)
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error finding book to update: %v", result.Error))
		return result.Error
	}

	book.Title = req.Title
	book.Author = req.Author
	book.Genre = req.Genre

	result = r.database.Save(&book)
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error updating book: %v", result.Error))
		return result.Error
	}

	return nil
}

func (r *bookRepository) Delete(id uint) error {
	var book *entity.Book

	result := r.database.First(&book, id)
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error finding book to delete: %v", result.Error))
		return result.Error
	}

	result = r.database.Delete(&book)
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error deleting book: %v", result.Error))
		return result.Error
	}

	return nil
}
