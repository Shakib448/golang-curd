package controllers

import (
	"errors"
	"net/http"

	"github.com/Shakib448/go-curd/initializers"
	"github.com/Shakib448/go-curd/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Sign_Up(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var singleUser models.User
	result_single := initializers.DB.Where(&models.User{Email: body.Email}).First(&singleUser)

	if result_single.Error != nil {
		if errors.Is(result_single.Error, gorm.ErrRecordNotFound) {
			hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Failed to hash password",
				})
				return
			}

			user := models.User{Email: body.Email, Password: string(hash)}

			initializers.DB.Create(&user)

			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exits",
		})
		return
	}

}
