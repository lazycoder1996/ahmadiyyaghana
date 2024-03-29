package controllers

import (
	"ahmadiyyaghana/initializers"
	"ahmadiyyaghana/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func CreateMember(c *gin.Context) {
	var body models.Member
	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res := initializers.DB.Create(&body)
	if res.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, res.Error.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"member": body,
	})

}
func UpdateMember(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Fullname, Hometown, Father, Mother, Position, Contact string
		Wassiyat, AimsCode                                    int
	}
	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var member models.Member
	updateBody := &models.Member{}
	deepcopier.Copy(body).To(updateBody)

	initializers.DB.First(&member, id)
	initializers.DB.Model(&member).UpdateColumns(&updateBody)

	c.IndentedJSON(http.StatusOK, member)

}
func DeleteMember(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Member{}, id)
	c.Status(http.StatusOK)
}

func GetMember(c *gin.Context) {
	aimsCode, err := strconv.ParseUint(c.Param("aimscode"), 10, 32)
	if err != nil {
		log.Fatalf("error")
	}
	var member models.Member
	initializers.DB.Where(&models.Member{AimsCode: uint(aimsCode)}).Find(&member)

	c.IndentedJSON(http.StatusOK, gin.H{
		"member": member,
	})
}

func GetMembers(c *gin.Context) {
	var members []models.Member
	initializers.DB.Find(&members)
	c.IndentedJSON(http.StatusOK, gin.H{
		"members": members,
	})
}
