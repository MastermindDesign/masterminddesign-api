package data

import (
	"context"
	"fmt"

	"github.com/FR0NK3NST33N/masterminddesign-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string
var Mongo *mongo.Client
var DB string

func init() {
	fmt.Println("Connecting to mongo...")
	var err error
	setEnv()
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(uri)
	Mongo, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
	}
}

func setEnv() {
	uri = utils.GoDotEnvVariable("MONGO_URI")
	DB = utils.GoDotEnvVariable("MONGO_DB")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	if DB == "" {
		DB = "mastermind"
	}
}
