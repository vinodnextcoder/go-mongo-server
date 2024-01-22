package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "context"
    // "time"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/vinodnextcoder/golang-mongo-server/mongoconnect"
    // "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/mongo/readpref"

    "github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vinodnextcoder/golang-mongo-server/docs"
  
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

  // Get Client, Context, CancelFunc and 
    // err from connect method.
    client, ctx, cancel, err := mongoconnect.Connect("mongodb://localhost:27017/test")
    if err != nil {
        panic(err)
    }
     
    // Release resource when the main
    // function is returned.
    defer mongoconnect.Close(client, ctx, cancel)
	
    // Ping mongoDB with Ping method
    mongoconnect.Ping(client, ctx)

    
    // Create  a object of type interface to  store
    // the bson values, that  we are inserting into database.
    var document interface{}
     
     
    document = bson.D{
        {"rollNo", 175},
        {"maths", 80},
        {"science", 90},
        {"computer", 95},
    }
     
    dbname := os.Getenv("DBNAME")
    // insertOne accepts client , context, database
    // name collection name and an interface that 
    // will be inserted into the  collection.
    // insertOne returns an error and a result of 
    // insert in a single document into the collection.
    insertOneResult, err := mongoconnect.Insertdata(client, ctx, dbname, "marks", document)
     
    // handle the error
    if err != nil {
        panic(err)
    }
     
    // print the insertion id of the document, 
    // if it is inserted.
    fmt.Println("Result of InsertOne")
    fmt.Println(insertOneResult.InsertedID)

  port := os.Getenv("PORT")

	router := gin.Default()

	router.GET("/", helloCall)
  router.POST("/add", mongoconnect.PostMarks)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	router.Run("0.0.0.0:"+port)
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
