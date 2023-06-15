package controllers

import (
	"net/http"

	"github.com/Shakib448/go-curd/initializers"
	"github.com/Shakib448/go-curd/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func Get_Post_By_Id(c *gin.Context) {

	var post models.Post
	if err := initializers.DB.First(&post, c.Param("id")).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Post not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve the post",
			})
		}
		return
	}

	// Response with the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}
