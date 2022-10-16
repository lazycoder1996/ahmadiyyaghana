package main

import (
	"ahmadiyyaghana/initializers"
	"ahmadiyyaghana/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Agent{}, &models.Admin{}, &models.Member{}, &models.User{})
}
