package dbs

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Act  string             `json:"act"`
	Done bool               `json:"done"`
}
