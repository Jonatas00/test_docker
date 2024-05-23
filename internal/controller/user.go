package controller

import (
	"test_docker/internal/database"
	"test_docker/internal/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	users := []models.User{}
	database.DB.Find(&users)
	ctx.JSON(200, users)
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	database.DB.Create(&user)
	ctx.JSON(200, user)
}

func GetUserByID(ctx *gin.Context) {
	var user models.User
	database.DB.Where("id = ?", ctx.Param("id")).First(&user)
	ctx.JSON(200, &user)
}

func UpdateUser(ctx *gin.Context) {
	var user models.User
	database.DB.Where("id = ?", ctx.Param("id")).First(&user)
	ctx.BindJSON(&user)
	database.DB.Save(&user)
	ctx.JSON(200, &user)
}

func DeleteUser(ctx *gin.Context) {
	var user models.User
	database.DB.Where("id = ?", ctx.Param("id")).Delete(&user)
	ctx.JSON(200, &user)
}
