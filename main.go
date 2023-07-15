package main

import (
	"log"

	"github.com/dasmino/curdgotest/controllers"
	"github.com/dasmino/curdgotest/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.POST("/posts", controllers.CreatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
