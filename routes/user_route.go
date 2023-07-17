package routes

import (
	"go_api_gin/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:_id", controllers.GetById())
	router.PUT("/user/:_id", controllers.EditUser())
	router.DELETE("/user/:_id", controllers.DeleteUser())
	router.GET("/users", controllers.GetAllUsers())
}
