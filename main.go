package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"context"
    "time"
 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"

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
    client, ctx, cancel, err := connect("mongodb://localhost:27017/test")
    if err != nil {
        panic(err)
    }
     
    // Release resource when the main
    // function is returned.
    defer close(client, ctx, cancel)
	
     
    // Ping mongoDB with Ping method
    ping(client, ctx)

  port := os.Getenv("PORT")

	router := gin.Default()

	router.GET("/", helloCall)
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
// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc){
	 
// CancelFunc to cancel to context
defer cancel()

// client provides a method to close 
// a mongoDB connection.
defer func(){

 // client.Disconnect method also has deadline.
 // returns error if any,
 if err := client.Disconnect(ctx); err != nil{
	 panic(err)
 }
}()
}

// This is a user defined method that returns mongo.Client, 
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and 
// resource associated with it.

func connect(uri string)(*mongo.Client, context.Context, 
				   context.CancelFunc, error) {
					
// ctx will be used to set deadline for process, here 
// deadline will of 30 seconds.
ctx, cancel := context.WithTimeout(context.Background(), 
								30 * time.Second)

// mongo.Connect return mongo.Client method
client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
return client, ctx, cancel, err
}

// This is a user defined method that accepts 
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error{

// mongo.Client has Ping to ping mongoDB, deadline of 
// the Ping method will be determined by cxt
// Ping method return error if any occurred, then
// the error can be handled.
if err := client.Ping(ctx, readpref.Primary()); err != nil {
 return err
}
fmt.Println("connected successfully")
return nil
}
