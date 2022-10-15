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
	initializers.DB.AutoMigrate(&models.Agent{})
	initializers.DB.AutoMigrate(&models.Member{})
}
