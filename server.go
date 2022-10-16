package main

import (
	"ahmadiyyaghana/controllers"
	"ahmadiyyaghana/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := gin.Default()
	api := r.Group("api")
	api.POST("/members", controllers.CreateMember)
	api.GET("/members", controllers.GetMembers)
	api.GET("/members/:id", controllers.GetMember)
	api.PUT("/members/:id", controllers.UpdateMember)
	api.DELETE("/members/:id", controllers.DeleteMember)
	r.Run(":" + port)
}
