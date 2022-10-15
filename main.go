package main

import (
	"ahmadiyyaghana/controllers"
	"ahmadiyyaghana/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	api := r.Group("api")
	api.POST("/agents", controllers.CreateAgent)
	api.GET("/agents", controllers.GetAgents)
	api.GET("/agents/:id", controllers.GetAgent)
	api.PUT("/agents/:id", controllers.UpdateAgent)
	api.DELETE("/agents/:id", controllers.DeleteAgent)

	api.POST("/members", controllers.CreateMember)
	api.GET("/members", controllers.GetMembers)
	api.GET("/members/:id", controllers.GetMember)
	api.PUT("/members/:id", controllers.UpdateMember)
	api.DELETE("/members/:id", controllers.DeleteMember)
	r.Run()
}
