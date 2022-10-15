package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name, Hometown, Father, Mother string `binding:"required"`
	Position, Contact              string
	AimsCode                       int `binding:"required"`
	Wassiyat                       int
}
