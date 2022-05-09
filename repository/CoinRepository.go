package repository

import (
	"context"
	"fmt"
	"klever-challenge/db"
	"klever-challenge/models"
	"strings"
	"time"

	bson "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const coinsCollection = "coins"

type CoinsRepository interface {
	GetAll() (*mongo.Cursor, error)
	Create(*models.Coin) error
	GetByName(string) (*mongo.SingleResult, error)
	UpdateCoin(bson.ObjectID, string) (*mongo.SingleResult, error)
	DeleteCoin(bson.ObjectID) (*mongo.SingleResult, error)
	AddUpvote(string) (*mongo.SingleResult, error)
	AddDownvote(string) (*mongo.SingleResult, error)
	RemoveUpvote(string) (*mongo.SingleResult, error)
	RemoveDownvote(string) (*mongo.SingleResult, error)
}

type coinRepository struct {
	db *mongo.Collection
}

func NewCoinRepository(conn db.Connection) CoinsRepository {
	return &coinRepository{db: conn.DB().Collection(coinsCollection)}
}

func (r *coinRepository) GetAll() (cursor *mongo.Cursor, err error) {
	// cur, err := r.db.Find(nil, nil)
	cursor, err = r.db.Find(context.Background(), bson.M{})

	return cursor, err
}

func (r *coinRepository) Create(coin *models.Coin) error {

	cursor, err := r.db.InsertOne(context.Background(), coin)
	fmt.Print(cursor)
	return err
}

func (r *coinRepository) GetByName(name string) (curso *mongo.SingleResult, err error) {
	name = strings.Title(name)
	cursor := r.db.FindOne(context.Background(), bson.M{"coinName": name})

	return cursor, err

}

func (r *coinRepository) UpdateCoin(id bson.ObjectID, name string) (cursor *mongo.SingleResult, err error) {
	name = strings.Title(name)
	Ctx := context.Background()
	filter := bson.M{"_id": id}

	cursor = r.db.FindOneAndUpdate(Ctx, filter, bson.M{"$set": bson.M{"coinName": name, "updatedAt": time.Now()}})

	return cursor, err
}

func (r *coinRepository) DeleteCoin(id bson.ObjectID) (cursor *mongo.SingleResult, err error) {
	fmt.Println(id)
	Ctx := context.Background()
	filter := bson.M{"_id": id}

	cursor = r.db.FindOneAndDelete(Ctx, filter)

	return cursor, err
}

func (r *coinRepository) AddUpvote(coinName string) (cursor *mongo.SingleResult, err error) {
	fmt.Println(coinName)
	Ctx := context.Background()
	filter := bson.M{"coinName": coinName}
	AddUpvote := bson.M{"$inc": bson.M{"upvote": 1}}

	cursor = r.db.FindOneAndUpdate(Ctx, filter, AddUpvote)

	return cursor, err
}

func (r *coinRepository) AddDownvote(coinName string) (cursor *mongo.SingleResult, err error) {
	fmt.Println(coinName)
	Ctx := context.Background()
	filter := bson.M{"coinName": coinName}
	AddDownvote := bson.M{"$inc": bson.M{"downvote": 1}}

	cursor = r.db.FindOneAndUpdate(Ctx, filter, AddDownvote)

	return cursor, err
}

func (r *coinRepository) RemoveUpvote(coinName string) (cursor *mongo.SingleResult, err error) {
	fmt.Println(coinName)
	Ctx := context.Background()
	filter := bson.M{"coinName": coinName}
	RemoveUpvote := bson.M{"$inc": bson.M{"upvote": -1}}

	cursor = r.db.FindOneAndUpdate(Ctx, filter, RemoveUpvote)

	return cursor, err
}

func (r *coinRepository) RemoveDownvote(coinName string) (cursor *mongo.SingleResult, err error) {
	fmt.Println(coinName)
	Ctx := context.Background()
	filter := bson.M{"coinName": coinName}
	RemoveDownvote := bson.M{"$inc": bson.M{"downvote": -1}}

	cursor = r.db.FindOneAndUpdate(Ctx, filter, RemoveDownvote)

	return cursor, err
}
