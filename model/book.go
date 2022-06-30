package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Author      string             `json:"author,omitempty" bson:"author,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}
