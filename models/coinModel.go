package models

import (
	"klever-challenge/app/pb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coin struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CoinName  string             `json:"coinName" bson:"coinName"`
	Upvote    int64              `json:"upvote" bson:"upvote"`
	Downvote  int64              `json:"downvote" bson:"downvote"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func (c *Coin) ToProtoBuffer() *pb.CoinResponse {
	return &pb.CoinResponse{
		Id:        c.ID.Hex(),
		CoinName:  c.CoinName,
		UpVotes:   c.Upvote,
		DownVotes: c.Downvote,
	}
}

func (c *Coin) ToProtoBufferListCoin() *pb.CoinsResponse {
	return &pb.CoinsResponse{
		Data: &pb.CoinData{
			Id:         c.ID.Hex(),
			CoinName:   c.CoinName,
			UpVotes:    c.Upvote,
			DownVotes:  c.Downvote,
			TotalVotes: 0,
			CreatedAt:  "",
			UpdatedAt:  "",
		},
	}
}
