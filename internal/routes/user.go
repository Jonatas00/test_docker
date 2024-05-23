package routes

import (
	"test_docker/internal/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/createuser", controller.CreateUser)
	router.GET("/getusers", controller.GetUsers)
	router.GET("/getuser/:id", controller.GetUserByID)
	router.PUT("/updateuser/:id", controller.UpdateUser)
	router.DELETE("deleteuser/:id", controller.DeleteUser)

}
