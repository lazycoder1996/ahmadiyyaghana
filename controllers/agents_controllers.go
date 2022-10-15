package controllers

import (
	"ahmadiyyaghana/initializers"
	"ahmadiyyaghana/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"gorm.io/gorm"
)

type ResBody struct {
	gorm.Model
	Fullname             string
	AimsCode             int
	Zone, Contact, Email string
}

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
	resBody := &ResBody{}
	deepcopier.Copy(body).To(resBody)
	c.IndentedJSON(http.StatusOK, gin.H{
		"user": resBody,
	})

}
func UpdateAgent(c *gin.Context) {
	id := c.Param("id")

	var body models.Agent
	c.Bind(&body)
	var user models.Agent
	initializers.DB.First(&user, id)
	initializers.DB.Model(&user).Updates(models.Agent{
		Fullname: body.Fullname,
		Contact:  body.Contact,
		Password: body.Password,
		Zone:     body.Zone,
	})
	resBody := &ResBody{}
	deepcopier.Copy(user).To(resBody)

	c.IndentedJSON(http.StatusOK, gin.H{
		"user": resBody,
	})

}
func DeleteAgent(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Agent{}, id)
	c.Status(http.StatusOK)
}
func GetAgent(c *gin.Context) {
	id := c.Param("id")
	var user models.Agent
	initializers.DB.First(&user, id)
	resBody := &ResBody{}
	deepcopier.Copy(user).To(resBody)

	c.IndentedJSON(http.StatusOK, gin.H{
		"user": resBody,
	})
}

func GetAgents(c *gin.Context) {
	var users []models.Agent
	initializers.DB.Find(&users)
	var resBody []ResBody
	for _, user := range users {
		rb := &ResBody{}
		deepcopier.Copy(user).To(rb)
		resBody = append(resBody, *rb)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"users": resBody,
	})
}
