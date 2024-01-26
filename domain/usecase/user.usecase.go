package usecase

import (
	"readmate/delivery/http/dto/request"
	"readmate/domain/repository"
)

type UserUsecase interface {
	Create(user request.RequestRegisterDTO) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (input *userUsecase) Create(user request.RequestRegisterDTO) error {
	if err := input.userRepository.Create(user); err != nil {
		return err
	}

	return nil
}
