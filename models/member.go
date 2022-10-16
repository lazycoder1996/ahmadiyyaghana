package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Fullname, Hometown, Father, Mother, Zone string `binding:"required"`
	Position, Contact                        string
	AimsCode                                 uint `gorm:"unique;autoIncrement:false" binding:"required"`
	Wassiyat, Role                           int
}
