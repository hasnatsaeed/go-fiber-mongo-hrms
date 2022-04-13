package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client

func MongoDbConnect(host, port, user, password, dbname string) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	client = dbClient

	log.Println("mongodb client connecting...")
	err = client.Connect(context.TODO())

	if err != nil {
		return nil, err
	}

	log.Println("mongodb client connected!")

	db := dbClient.Database(dbname)

	return db, nil
}

func MongoDbDisconnect() error {
	log.Println("mongodb client disconnecting...")
	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	log.Println("mongodb client disconnected")
	return nil

}
