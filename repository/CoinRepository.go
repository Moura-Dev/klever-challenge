package repository

import (
	"context"
	"fmt"
	"klever-challenge/models"

	bson "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CoinsRepositoryInterface interface {
	GetAll() (mongo.Cursor, error)
	GetById(id string) (*models.Coin, error)
	GetByName(name string) (*models.Coin, error)
	Create(coin *models.Coin) error
	Update(coin *models.Coin) error
	DropCollection() error
	Delete(id string) error
}

type CoinsRepository struct {
	Db  *mongo.Collection
	Ctx context.Context
}

// GetAll implements CoinsRepository
func (r *CoinsRepository) GetAll() (cur mongo.Cursor, err error) {
	cursor, err := r.Db.Find(r.Ctx, bson.M{}, options.Find().SetSort(bson.M{}))
	if err != nil {
		return cur, status.Errorf(codes.Internal, fmt.Sprintf("Not found %v", err))
	}
	defer cursor.Close(r.Ctx)

	return cur, nil
}

// GetByName implements CoinsRepository
func (r *CoinsRepository) GetByName(name string) (coin *models.Coin, err error) {
	cursor, err := r.Db.Find(r.Ctx, bson.M{"name": name}, options.Find().SetSort(bson.M{}))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Not found %v", err))
	}
	cursor.All(r.Ctx, &coin)
	return coin, nil
}

// GetById implements CoinsRepository
func (r *CoinsRepository) GetById(id string) (coin *models.Coin, err error) {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid id: %v", err))
	}
	cursor, err := r.Db.Find(r.Ctx, bson.M{"_id": objectId}, options.Find().SetSort(bson.M{}))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Not found %v", err))
	}
	cursor.All(r.Ctx, &coin)
	return coin, nil
}

// CreateCoin implements CoinsRepository
func (r *CoinsRepository) Create(coin *models.Coin) error {
	_, err := r.Db.InsertOne(r.Ctx, []interface{}{coin})

	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Insert Failed error: %v", err))
	}
	return nil
}

// Update implements CoinsRepository
func (r *CoinsRepository) Update(coin *models.Coin) error {

	_, err := r.Db.UpdateByID(r.Ctx, coin.ID, coin)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Update Failed error: %v", err))
	}
	return nil
}

func (r *CoinsRepository) DropCollection() error {
	err := r.Db.Drop(r.Ctx)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Drop Collection Failed error: %v", err))
	}
	return nil
}

func (r *CoinsRepository) Delete(id string) error {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid id: %v", err))

	}
	cursor, err := r.Db.DeleteOne(r.Ctx, bson.M{"_id": objectId})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Delete Failed error: %v", err))
	}
	return cursor
}
