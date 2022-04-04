package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Employee struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}

var client *mongo.Client
var db *mongo.Database

func Connect() {
	uri := "mongodb://root:jenkins@127.0.0.1:27017"
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	client = dbClient

	log.Println("connecting mongodb client...")
	err = client.Connect(context.TODO())

	if err != nil {
		panic(err)
	}

	log.Println("mongodb client connected")

	db = dbClient.Database("go-fiber-mongo-hrms")
}

func Disconnect() {
	log.Println("disconnecting mongodb client...")
	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	log.Println("mongodb client disconnected")

}

func GetCollection(collection string) *mongo.Collection {
	return db.Collection(collection)
}
