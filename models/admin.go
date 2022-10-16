package models

type Admin struct {
	Info Member `gorm:"embedded" binding:"required"`
}
