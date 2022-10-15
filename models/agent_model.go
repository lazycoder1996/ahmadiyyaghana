package models

import "gorm.io/gorm"

type Agent struct {
	gorm.Model
	Fullname             string `binding:"required"`
	Password             string `binding:"required"`
	AimsCode             int    `binding:"required"`
	Zone, Contact, Email string `binding:"required"`
}
