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

var ctx context.Context

func init() {
	ctx = context.TODO()
}
func Connect() {
	uri := "mongodb://root:jenkins@mongodb:27017"
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	client = dbClient

	log.Println("connecting mongodb client...")
	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	log.Println("mongodb client connected")

	db = dbClient.Database("go-fiber-mongo-hrms")
}

func Disconnect() {
	log.Println("disconnecting mongodb client...")
	err := client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	log.Println("mongodb client disconnected")

}

func GetDb() *mongo.Database {
	return db
}

func GetCtx() *context.Context {
	return &ctx
}
