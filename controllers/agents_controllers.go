package controllers

import (
	"ahmadiyyaghana/initializers"
	"ahmadiyyaghana/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func CreateAgent(c *gin.Context) {
	var body models.Agent
	if err := c.Bind(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return

	}
	res := initializers.DB.Create(&body)
	if res.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"agent": body,
	})

}
func UpdateAgent(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Fullname             string
		Password             string
		AimsCode             int
		Zone, Contact, Email string
	}

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	updateBody := &models.Agent{}
	deepcopier.Copy(body).To(updateBody)

	var agent models.Agent

	initializers.DB.First(&agent, id)
	initializers.DB.Model(&agent).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, gin.H{
		"agent": agent,
	})

}
func DeleteAgent(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Agent{}, id)
	c.Status(http.StatusOK)
}
func GetAgent(c *gin.Context) {
	id := c.Param("id")
	var agent models.Agent
	initializers.DB.First(&agent, id)

	c.IndentedJSON(http.StatusOK, gin.H{
		"agent": agent,
	})
}

func GetAgents(c *gin.Context) {
	var agents []models.Agent
	initializers.DB.Find(&agents)
	c.IndentedJSON(http.StatusOK, gin.H{
		"agents": agents,
	})
}
