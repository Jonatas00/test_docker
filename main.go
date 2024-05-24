package main

import (
	"test_docker/internal/config"
	"test_docker/internal/database"
	"test_docker/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	router := gin.New()

	database.Connect()
	routes.UserRoute(router)

	router.Run(":" + config.APP_PORT)
}
