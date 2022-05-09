// The Technical Challenge consists of creating an API with Golang using
// gRPC with stream pipes that exposes an Upvote service endpoints.
//  The API will provide the user an interface to upvote or downvote a
// known list of the main Cryptocurrencies (Bitcoin, ethereum, litecoin, etc..).

// Technical requirements:

// Keep the code in Github
// The API must have a read, insert, delete and update interfaces.
// The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string.
// The API must contain unit test of methods it uses
// You can choose the database but the structs used with it should support Marshal/Unmarshal with bson, json and struct

// Extra:

// Deliver the whole solution running in some free cloud service
// The API have a method that stream a live update of the current sum of the votes from a given Cryptocurrency

package main

import (
	"fmt"
	"klever-challenge/app/pb"
	"klever-challenge/controllers"
	"klever-challenge/db"
	"klever-challenge/repository"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	coinsRepository := repository.NewCoinRepository(conn)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "50051"))
	if err != nil {
		log.Fatalf("Failed connect on port :%s: %s", "50051", err.Error())

	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	upvoteService := controllers.UpvoteServiceServer{
		CoinsRepository: coinsRepository,
	}

	pb.RegisterUpvoteServiceServer(grpcServer, &upvoteService)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to connect serve: %v", err)
		}
	}()
	fmt.Printf("Server started on port :%s\n", "50051")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	grpcServer.Stop()
	listener.Close()
	conn.Close()
	fmt.Println("Server stopped.")
}
