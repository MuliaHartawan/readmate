package entity

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID     uint
	BookID     uint
	ReviewText string
}
