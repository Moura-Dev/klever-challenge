package repository

import (
	"context"
	"fmt"
	"klever-challenge/db"
	"klever-challenge/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const coinsCollection = "coins"

type CoinsRepository interface {
	GetAll() (*mongo.Cursor, error)
	Create(*models.Coin) error
}

type coinRepository struct {
	db *mongo.Collection
}

func NewCoinRepository(conn db.Connection) CoinsRepository {
	return &coinRepository{db: conn.DB().Collection(coinsCollection)}
}

func (r *coinRepository) GetAll() (curso *mongo.Cursor, err error) {
	// cur, err := r.db.Find(nil, nil)
	cursor, err := r.db.Find(context.Background(), bson.M{}, options.Find().SetSort(bson.M{}))

	if err != nil {
		return nil, err
	}
	fmt.Println(cursor)
	return curso, nil
}

func (r *coinRepository) Create(coin *models.Coin) error {

	cursor, err := r.db.InsertOne(context.Background(), coin)
	fmt.Println(cursor, err)
	return nil
}
