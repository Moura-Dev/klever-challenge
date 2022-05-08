package controllers

import (
	"context"
	"fmt"
	"klever-challenge/app/pb"
	"klever-challenge/models"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpvoteServiceServer struct {
	Db  *mongo.Collection
	Ctx context.Context
	pb.UnimplementedUpvoteServiceServer
}

func (s *UpvoteServiceServer) ListCoin(req *pb.ListCoinsRequest, stream pb.UpvoteService_ListCoinsServer) error {

	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Not found"))
	}
	coins := models.Coin{}

	for cursor.Next(s.Ctx) {
		err := cursor.Decode(&coins)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		stream.Send(&pb.CoinsResponse{
			Data: &pb.CoinData{},
		})
	}
	return nil
}
