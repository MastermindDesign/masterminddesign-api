package project

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Stack string             `json:"stack,omitempty" bson:"stack,omitempty"`
	Link  string             `json:"link,omitempty" bson:"link,omitempty"`
	Order int                `json:"order,omitempty" bson:"order,omitempty"`
}
