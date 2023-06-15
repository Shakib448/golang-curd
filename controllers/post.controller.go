package controllers

import (
	"github.com/Shakib448/go-curd/initializers"
	"github.com/Shakib448/go-curd/models"
	"github.com/gin-gonic/gin"
)

func Post_Create(c *gin.Context) {
	// Get data of request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func Get_Posts(c *gin.Context) {
	// Get the posts
	var getPosts []models.Post
	initializers.DB.Find(&getPosts)

	// Response with them
	c.JSON(200, gin.H{
		"posts": getPosts,
	})
}
