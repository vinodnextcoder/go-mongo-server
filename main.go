package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vinodnextcoder/golang-mongo-server/docs"
	"github.com/vinodnextcoder/golang-mongo-server/routes"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample gin web server

// @contact.name   vinod
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:3001
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	err := godotenv.Load(".env.dev")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//run database
	// configs.ConnectDB()

	port := os.Getenv("PORT")

	router := gin.Default()

	router.GET("/", helloCall)
	routes.UserRoute(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("0.0.0.0:" + port)
}

// helloCall godoc
// @Summary hellow example
// @Schemes
// @Description Hello
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Hello, You created a Web App!
// @Router / [get]
func helloCall(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, You created a Web App!"})
}
