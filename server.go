package main

import (
	"ahmadiyyaghana/controllers"
	"ahmadiyyaghana/initializers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.LoadHTMLGlob("templates/*.tmpl.html")
	r.Static("/static", "static")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl.html", nil)

	})
	api := r.Group("api")
	api.POST("/members", controllers.CreateMember)
	api.GET("/members", controllers.GetMembers)
	api.GET("/members/:aimscode", controllers.GetMember)
	api.PUT("/members/:id", controllers.UpdateMember)
	api.DELETE("/members/:id", controllers.DeleteMember)
	r.Run(":" + port)
}
