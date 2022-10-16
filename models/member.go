package models

type Member struct {
	ID                                       uint   `gorm:"primaryKey"`
	Fullname, Hometown, Father, Mother, Zone string `binding:"required"`
	Position, Contact                        string
	AimsCode                                 uint `gorm:"unique;autoIncrement:false" binding:"required"`
	Wassiyat, Role                           int
}
