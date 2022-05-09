package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type conn struct {
	session  *mongo.Client
	database *mongo.Database
}

func Connect() (Connection, error) {
	// Set client options
	mongoCtx := context.Background()
	// Create a new client and connect to the server
	client, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	return &conn{session: client, database: client.Database("klever-challenge")}, nil
}

func (c *conn) Close() {
	c.session.Disconnect(context.Background())
}

func (c *conn) DB() *mongo.Database {
	return c.database
}
