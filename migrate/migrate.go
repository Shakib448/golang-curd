package main

import (
	"github.com/Shakib448/go-curd/initializers"
	"github.com/Shakib448/go-curd/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
