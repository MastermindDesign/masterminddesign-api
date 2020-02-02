package project

import (
	"context"
	"log"

	"github.com/FR0NK3NST33N/masterminddesign-api/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAll() []Project {
	ctx := context.Background()
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"order", 1}})
	var results []Project
	coll := data.Mongo.Database(data.DB).Collection("projects")
	cursor, err := coll.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var project Project
		err := cursor.Decode(&project)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, project)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(ctx)

	return results
}
