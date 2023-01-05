package dbs

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Uri            string = "mongodb://localhost:27017/"
	dbName         string = "reactTodo"
	collectionName string = "reactTodo"
	TodoCollection *mongo.Collection
)

func init() {
	var clientOption = options.Client().ApplyURI(Uri)
	var connection, connectionErr = mongo.Connect(context.Background(), clientOption)

	if connectionErr != nil {
		log.Println(connectionErr.Error())
		os.Exit(2)
	}

	TodoCollection = connection.Database(dbName).Collection(collectionName)
}
