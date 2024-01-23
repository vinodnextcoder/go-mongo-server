package routes

import (
	"github.com/gin-gonic/gin"
	markscontrollers "github.com/vinodnextcoder/golang-mongo-server/controllers"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", markscontrollers.CreateUser())
}
