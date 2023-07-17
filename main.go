package main

import (
	// "errors"
	// "net/http"

	"go_api_gin/configs"
	"go_api_gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()

	router := gin.Default()
	routes.UserRoute(router)
	router.Run(configs.EnvPort())
}
