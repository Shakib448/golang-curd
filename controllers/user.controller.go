package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Shakib448/go-curd/initializers"
	"github.com/Shakib448/go-curd/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func Sign_In(c *gin.Context) {

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

	var user models.User
	initializers.DB.Where(&models.User{Email: body.Email}).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to creating token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
	})
}
