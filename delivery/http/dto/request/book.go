package request

import "readmate/domain/entity"

type RequestBookCreateDTO struct {
	Title  string `json:"password"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}
type RequestBookUpdateDTO struct {
	Title  string `json:"password"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

func (input RequestBookCreateDTO) Create() entity.Book {
	return entity.Book{
		Title:  input.Title,
		Author: input.Author,
		Genre:  input.Genre,
	}
}

func (input RequestBookCreateDTO) Update() entity.Book {
	return entity.Book{
		Title:  input.Title,
		Author: input.Author,
		Genre:  input.Genre,
	}
}
