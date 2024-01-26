package domain

import (
	"readmate/domain/repository"
	"readmate/domain/usecase"
	"readmate/infrastructure"
)

type Domain struct {
	User usecase.UserUsecase
}

func DomainInit() Domain {

	database, _ := infrastructure.DatabaseInit()

	userRepository := repository.NewUserRepository(database)

	userUsecase := usecase.NewUserUsecase(userRepository)

	return Domain{
		User: userUsecase,
	}
}
