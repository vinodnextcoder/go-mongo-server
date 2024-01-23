package routes

import (
	"github.com/gin-gonic/gin"
	markscontrollers "github.com/vinodnextcoder/golang-mongo-server/controllers"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", markscontrollers.CreateUser())
	router.GET("/user/:userId", markscontrollers.GetAUser())
	router.PUT("/user/:userId", markscontrollers.EditAUser())
	router.DELETE("/user/:userId", markscontrollers.DeleteAUser())
	router.GET("/users", markscontrollers.GetAllUsers())
}
