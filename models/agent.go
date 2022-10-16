package models

type Agent struct {
	Info Member `gorm:"embedded" binding:"required"`
}
