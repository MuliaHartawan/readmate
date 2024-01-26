package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	UserID uint
	Title  string
	Author string
	Rating int
}
