package main

import (
	"context"
	"derek-api/controllers"
	"derek-api/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server 			*gin.Engine
	userservice 	services.UserService
	usercontroller 	controllers.UserController
	ctx 			context.Context
	usercollection 	*mongo.Collection
	mongoclient 	*mongo.Client
	err 			error
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

func init() {
	ctx = context.TODO()	// todo will create a 'Context Object'

	dotenv := goDotEnvVariable("MONGO_URL")	// get env variable

	// mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")		// connection logic
	mongoconn := options.Client().ApplyURI(dotenv)		// connection logic

	mongoclient, err = mongo.Connect(ctx, mongoconn)	// initialize mongo client

	if err != nil {
		log.Fatal(err)	// close app
	}

	err = mongoclient.Ping(ctx, readpref.Primary())	// ping mongodb to check error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection has been established")

	// Initializations
	usercollection =  mongoclient.Database("userdb").Collection("users") // Initialize User Collection
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)	 // disconnect from mongo if app shut downs
	
	basepath := server.Group("/v1")	// v1/user/...

	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}
