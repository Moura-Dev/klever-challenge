package controllers

import (
	"context"
	"fmt"
	"klever-challenge/app/pb"
	"klever-challenge/models"
	"klever-challenge/repository"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	bson "go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpvoteServiceServer struct {
	repository.CoinsRepository
	pb.UnimplementedUpvoteServiceServer
}

func (s *UpvoteServiceServer) ListCoins(req *pb.ListCoinsRequest, stream pb.UpvoteService_ListCoinsServer) error {
	cursor, err := s.CoinsRepository.GetAll()
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	var coin models.Coin
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&coin)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		stream.Send(coin.ToProtoBufferListCoin())
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

func (s *UpvoteServiceServer) GetCoinByName(ctx context.Context, req *pb.CoinNameRequest) (*pb.CoinResponse, error) {
	name := strings.Title(req.GetCoinName())
	cursor, err := s.CoinsRepository.GetByName(name)
	if err != nil {
		return nil, err
	}
	var coin models.Coin
	cursor.Decode(&coin)
	return coin.ToProtoBuffer(), nil
}

func (s *UpvoteServiceServer) CreateCoin(ctx context.Context, req *pb.CoinNameRequest) (*pb.CoinResponse, error) {

	coin := models.Coin{
		ID:        primitive.NewObjectID(),
		CoinName:  strings.Title(req.GetCoinName()),
		Upvote:    0,
		Downvote:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.CoinsRepository.Create(&coin)
	if err != nil {
		return nil, err
	}
	coin.ToProtoBuffer()
	fmt.Print(coin)
	return coin.ToProtoBuffer(), nil

}

func (s *UpvoteServiceServer) UpdateCoin(ctx context.Context, req *pb.UpdateCoinRequest) (*pb.CoinResponse, error) {
	objectId, err := bson.ObjectIDFromHex(req.GetId())
	name := strings.Title(req.GetCoinName())
	cursor, err := s.CoinsRepository.UpdateCoin(objectId, name)
	if err != nil {
		return nil, err
	}
	var coin models.Coin

	coin.CoinName = req.CoinName
	coin.ID = objectId
	fmt.Println(&cursor)
	return coin.ToProtoBuffer(), nil

}

func (s *UpvoteServiceServer) RemoveCoin(ctx context.Context, req *pb.CoinIdRequest) (*pb.CoinResponse, error) {
	objectId, err := bson.ObjectIDFromHex(req.GetId())
	cursor, err := s.CoinsRepository.DeleteCoin(objectId)
	if err != nil {
		return nil, err
	}
	var coin models.Coin
	coin.ID = objectId
	fmt.Println(&cursor)
	return coin.ToProtoBuffer(), nil

}

func (s *UpvoteServiceServer) UpVote(ctx context.Context, req *pb.CoinNameRequest) (*pb.CoinResponse, error) {
	coinName := strings.Title(req.GetCoinName())

	cursor, err := s.CoinsRepository.AddUpvote(coinName)
	if err != nil {
		return nil, err
	}

	if cursor == nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Coin Not found in database %s", req.GetCoinName()),
		)
	}

	var coin models.Coin
	coin.CoinName = coinName
	fmt.Println(&cursor)
	return coin.ToProtoBuffer(), nil

}

func (s *UpvoteServiceServer) DownVote(ctx context.Context, req *pb.CoinNameRequest) (*pb.CoinResponse, error) {
	coinName := strings.Title(req.GetCoinName())

	cursor, err := s.CoinsRepository.AddDownvote(coinName)
	if err != nil {
		return nil, err
	}

	if cursor == nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Coin Not found in database %s", req.GetCoinName()),
		)
	}

	var coin models.Coin
	coin.CoinName = coinName
	fmt.Println(&cursor)
	return coin.ToProtoBuffer(), nil

}

func (s *UpvoteServiceServer) RemoveUpVote(ctx context.Context, req *pb.CoinNameRequest) (*pb.CoinResponse, error) {
	coinName := strings.Title(req.GetCoinName())

	cursor, err := s.CoinsRepository.RemoveUpvote(coinName)
	if err != nil {
		return nil, err
	}

	if cursor == nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Coin Not found in database %s", req.GetCoinName()),
		)
	}

	var coin models.Coin
	coin.CoinName = coinName
	fmt.Println(&cursor)
	return coin.ToProtoBuffer(), nil

}

func (s *UpvoteServiceServer) RemoveDownvote(ctx context.Context, req *pb.CoinNameRequest) (*pb.CoinResponse, error) {
	coinName := strings.Title(req.GetCoinName())

	cursor, err := s.CoinsRepository.RemoveDownvote(coinName)
	if err != nil {
		return nil, err
	}

	if cursor == nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Coin Not found in database %s", req.GetCoinName()),
		)
	}

	var coin models.Coin
	coin.CoinName = coinName
	fmt.Println(&cursor)
	return coin.ToProtoBuffer(), nil

}
