package models

type User struct {
	Info Member `gorm:"embedded" binding:"required"`
}
