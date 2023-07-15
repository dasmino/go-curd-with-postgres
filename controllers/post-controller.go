package controllers

import (
	"fmt"

	"github.com/dasmino/curdgotest/database"
	"github.com/dasmino/curdgotest/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	DB := database.ConnectToDB()
	DB.Find(&posts)
	c.JSON(200, gin.H{
		"message": "Success",
		"posts":   posts,
	})
}

func CreatePost(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	post := &models.Post{Title: body.Title, Body: body.Body}
	DB := database.ConnectToDB()
	result := DB.Create(&post)
	if result.Error != nil {
		fmt.Println(result)
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
		"posts":   post,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	DB := database.ConnectToDB()
	DB.First(&post, id)

	c.JSON(200, gin.H{
		"message": "Success",
		"posts":   post,
	})
}

func UpdatePost(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}
	id := c.Param("id")
	c.Bind(&body)
	var post models.Post
	DB := database.ConnectToDB()
	DB.First(&post, id)

	DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"message": "Success",
		"posts":   post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	DB := database.ConnectToDB()
	DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "Success",
	})
}
