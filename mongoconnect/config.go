package mongoconnect

import (
	"fmt"
	"context"
    "time"
    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)


// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func Close(client *mongo.Client, ctx context.Context,
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

func Connect(uri string)(*mongo.Client, context.Context, 
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
func Ping(client *mongo.Client, ctx context.Context) error{

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

type marks struct {
    rollNo     int  `json:"rollno"`
    maths  int  `json:"maths"`
    science int  `json:"science"`
    computer  int `json:"computer"`
}

//  for testing added

func PostMarks(c *gin.Context) {
    var newMarks marks

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newMarks); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newMarks)
    c.IndentedJSON(http.StatusCreated, newMarks)
}