package main

import (
	"github.com/Shakib448/go-curd/controllers"
	"github.com/Shakib448/go-curd/initializers"
	"github.com/Shakib448/go-curd/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/post-create", controllers.Post_Create)
	r.GET("/post-get", controllers.Get_Posts)
	r.GET("/post-get/:id", controllers.Get_Post_By_Id)
	r.PUT("/post-update/:id", controllers.Post_Update)
	r.DELETE("/post-delete/:id", controllers.Post_Delete)

	// For users
	r.POST("/create-user", controllers.Sign_Up)
	r.POST("/sign-in", controllers.Sign_In)
	r.GET("/validate", middleware.Auth, controllers.Validate)
	r.Run()

}
