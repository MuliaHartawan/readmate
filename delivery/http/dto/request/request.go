package request

import "readmate/domain/entity"

type RequestRegisterDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Phone    int    `json:"phone"`
}

type RequestLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (input RequestRegisterDTO) Register() entity.User {
	return entity.User{
		Email:    input.Email,
		Password: input.Password,
		Username: input.Username,
		Phone:    input.Phone,
	}
}

func (input RequestLoginDTO) Login() entity.User {
	return entity.User{
		Email:    input.Email,
		Password: input.Password,
	}
}
