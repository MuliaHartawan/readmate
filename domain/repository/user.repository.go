package repository

import (
	"fmt"
	"log"
	"readmate/delivery/http/dto/request"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(value request.RequestRegisterDTO) error
}

type databaseRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) databaseRepository {
	return databaseRepository{
		database: database,
	}
}

func (input databaseRepository) Create(value request.RequestRegisterDTO) error {
	result := input.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error create user:: %v", result.Error))
		return result.Error
	}

	return nil
}
