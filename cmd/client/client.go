package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/linielson/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()
	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Lini",
		Email: "linielsonrosa@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Linielson",
		Email: "linielson.rosa@gmail.com",
	}

	resStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}

		fmt.Println("Status:", stream.Status, "-", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "l1",
			Name:  "Lini 1",
			Email: "l@l1.com",
		},
		&pb.User{
			Id:    "l2",
			Name:  "Lini 2",
			Email: "l@l2.com",
		},
		&pb.User{
			Id:    "l3",
			Name:  "Lini 3",
			Email: "l@l3.com",
		},
		&pb.User{
			Id:    "l4",
			Name:  "Lini 4",
			Email: "l@l4.com",
		},
		&pb.User{
			Id:    "l5",
			Name:  "Lini 5",
			Email: "l@l5.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
