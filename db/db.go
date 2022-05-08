package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// URL DB
const uri = "mongodb://localhost:27017"

func Connect() (*mongo.Collection, context.Context) {

	mongoCtx := context.Background()

	// Create client and connect
	client, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Success connected.")

	// DB collection
	kleverDB := client.Database("klever-challenge").Collection("coins")

	return kleverDB, mongoCtx
}
