package usecase

import (
	"errors"
	"fmt"
	"log"
	"readmate/delivery/http/dto/request"
	"readmate/domain/entity"
	"readmate/domain/repository"
)

type BookUsecase interface {
	Create(book *request.RequestBookCreateDTO) error
	GetAll() ([]*entity.Book, error)
	GetByID(id uint) (*entity.Book, error)
	Update(id uint, book *request.RequestBookUpdateDTO) error
	Delete(id uint) error
}

type bookUsecase struct {
	bookRepository repository.BookRepository
}

func NewBookUsecase(bookRepository repository.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepository: bookRepository,
	}
}

func (input *bookUsecase) Create(book *request.RequestBookCreateDTO) error {
	if err := input.bookRepository.Create(book); err != nil {
		return err
	}

	return nil
}

func (input *bookUsecase) GetAll() ([]*entity.Book, error) {
	books, err := input.bookRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (input *bookUsecase) GetByID(id uint) (*entity.Book, error) {
	book, err := input.bookRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if book == nil {
		log.Println(fmt.Sprintf("Book with ID: %v not found", id))
		return nil, errors.New(fmt.Sprintf("Book with ID: %v not found", id))
	}

	return book, nil
}

func (input *bookUsecase) Update(id uint, req *request.RequestBookUpdateDTO) error {

	book, err := input.bookRepository.GetByID(id)
	if err != nil {
		return err
	}

	if book == nil {
		log.Println(fmt.Sprintf("Book with ID: %v not found", id))
		return errors.New(fmt.Sprintf("Book with ID: %v not found", id))
	}

	if err := input.bookRepository.Update(id, req); err != nil {
		return err
	}

	return nil
}

func (input *bookUsecase) Delete(id uint) error {
	book, err := input.bookRepository.GetByID(id)
	if err != nil {
		return err
	}

	if book == nil {
		log.Println(fmt.Sprintf("Book with ID: %v not found", id))
		return errors.New(fmt.Sprintf("Book with ID: %v not found", id))
	}

	if err := input.bookRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
