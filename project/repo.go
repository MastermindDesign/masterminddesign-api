package project

import (
	"context"
	"log"

	"github.com/FR0NK3NST33N/masterminddesign-api/data"
	"go.mongodb.org/mongo-driver/bson"
)

func FindAll() []Project {
	ctx := context.Background()
	var results []Project
	coll := data.Mongo.Database(data.DB).Collection("projects")
	cursor, err := coll.Find(ctx, bson.M{})
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
