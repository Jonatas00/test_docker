package main

import (
	"test_docker/internal/database"
	"test_docker/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	database.Connect()
	routes.UserRoute(router)

	router.Run()
}
