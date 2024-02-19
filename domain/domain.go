package domain

import (
	"readmate/domain/repository"
	"readmate/domain/usecase"
	"readmate/infrastructure"
)

type Domain struct {
	User usecase.UserUsecase
	Book usecase.BookUsecase
}

func DomainInit() Domain {

	database, _ := infrastructure.DatabaseInit()

	userRepository := repository.NewUserRepository(database)
	bookRepositroy := repository.NewBookRepository(database)

	userUsecase := usecase.NewUserUsecase(userRepository)
	bookUsecase := usecase.NewBookUsecase(bookRepositroy)

	return Domain{
		User: userUsecase,
		Book: bookUsecase,
	}
}
