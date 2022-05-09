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
)

type UpvoteServiceServer struct {
	repository.CoinsRepository
	pb.UnimplementedUpvoteServiceServer
}

func (s *UpvoteServiceServer) ListCoins(req *pb.ListCoinsRequest, stream pb.UpvoteService_ListCoinsServer) error {
	coins, err := s.CoinsRepository.GetAll()

	fmt.Println(coins, err)
	return nil
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

	// fmt.Println(coin, err)
	fmt.Print(coin)

	return nil, err
}
