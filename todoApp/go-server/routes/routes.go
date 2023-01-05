package routes

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/praveenmahasena647/todo/dbs"
)

var handlecors = cors.New(cors.Config{
	AllowOrigins:  []string{"*"},
	AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
	AllowHeaders:  []string{"Origin", "Content-Type"},
	ExposeHeaders: []string{"Content-Length", "Content-Type"},
})

func RunServer() *gin.Engine {
	var router *gin.Engine = gin.Default()
	router.Use(handlecors)

	router.GET("/", getAll)
	router.GET("/:id", getOne)
	router.POST("/", postOne)
	router.PUT("/:id", putOne)
	router.DELETE("/:id", deleteOne)

	return router
}

func getAll(c *gin.Context) {
	var todos []dbs.ToDo
	var cursor, err = dbs.TodoCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSONP(400, "dbError")
		return
	}

	for cursor.Next(context.TODO()) {
		var todo dbs.ToDo
		var err = cursor.Decode(&todo)
		if err != nil {
			c.JSONP(400, "dbError")
			return
		}
		todos = append(todos, todo)
	}

	c.JSONP(200, todos)
}

func getOne(c *gin.Context) {
	var id = c.Param("id")
	var ID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSONP(400, "mongoId decode error")
		return
	}
	var todo *dbs.ToDo
	dbs.TodoCollection.FindOne(context.TODO(), bson.M{"_id": ID}).Decode(&todo)

	c.JSONP(200, todo)
}

func postOne(c *gin.Context) {
	var todo *dbs.ToDo = &dbs.ToDo{}
	json.NewDecoder(c.Request.Body).Decode(&todo)

	var _, doneErr = dbs.TodoCollection.InsertOne(context.TODO(), todo)
	if doneErr != nil {
		log.Println(doneErr.Error())
		return
	}
	c.Status(200)
}

func putOne(c *gin.Context) {
	var id = c.Param("id")
	var Id, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Status(400)
		return
	}

	var todo *dbs.ToDo = &dbs.ToDo{}
	json.NewDecoder(c.Request.Body).Decode(&todo)
	var done, doneErr = dbs.TodoCollection.UpdateOne(context.TODO(), bson.M{"_id": Id}, bson.M{"$set": bson.A{todo}})
	log.Println(done, doneErr)
	c.Status(200)
}

func deleteOne(c *gin.Context) {
	var id = c.Param("id")
	var ID, idErr = primitive.ObjectIDFromHex(id)
	if idErr != nil {
		c.Status(400)
		return
	}
	var done = dbs.TodoCollection.FindOneAndDelete(context.TODO(), bson.M{"_id": ID})
	println(done)
}
