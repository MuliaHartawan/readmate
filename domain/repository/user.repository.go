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

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *userRepository {
	return &userRepository{
		database: database,
	}
}

func (input userRepository) Create(value request.RequestRegisterDTO) error {
	result := input.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error create user:: %v", result.Error))
		return result.Error
	}

	return nil
}
