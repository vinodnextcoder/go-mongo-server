package mongoconnect

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type marksData struct {
	rollNo   int
	maths    int
	science  int
	computer int
}

// albums slice to seed record album data.
var usermarks = []marksData{
	{rollNo: 2, maths: 2, science: 2, computer: 56},
	{rollNo: 32, maths: 2, science: 2, computer: 17},
	{rollNo: 3, maths: 1, science: 2, computer: 39},
}

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
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

func Connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func Ping(client *mongo.Client, ctx context.Context) error {

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

// insertOne is a user defined method, used to insert
// documents into collection returns result of InsertOne
// and error if any.
func Insertdata(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func Insertdata1(dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	// connect mongodb and insert records
	client, ctx, cancel, err := Connect("mongodb://localhost:27017/test")
	if err != nil {
		panic(err)
	}

	// select database and collection ith Client.Database method
	// and Database.Collection method
	fmt.Println("Result of InsertOne", cancel)
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

//  for testing added

func PostMarks(c *gin.Context) {

	// Get the request body

	// Decode the request body into a JSON object
	var jsonBody map[string]interface{}
	err := c.BindJSON(&jsonBody)
	if err != nil {
		// Handle error
		return
	}
	Connect("mongodb://localhost:27017/test")
	dbname := os.Getenv("DBNAME")
	// var client mongo.Client
	// var ctx context.Context
	// insertOne accepts client , context, database
	// name collection name and an interface that
	// will be inserted into the  collection.
	// insertOne returns an error and a result of
	// insert in a single document into the collection.

	insertOneResult, err := Insertdata1(dbname, "marks", jsonBody)

	fmt.Println(err)

	// handle the error
	if err != nil {
		panic(err)
	}

	// print the insertion id of the document,
	// if it is inserted.
	fmt.Println("Result of InsertOne")
	fmt.Println(insertOneResult.InsertedID)
	// Use the JSON object
	fmt.Println(jsonBody)
	c.IndentedJSON(http.StatusCreated, jsonBody)
}
