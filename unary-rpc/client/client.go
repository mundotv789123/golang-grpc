package client

import (
	"context"
	"fmt"
	"unary-rpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	dial, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	userClient := pb.NewUserServiceClient(dial)
	user, err := userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "1",
		Age:  40,
		Name: "Teste",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Primeiro usuário criado: %v\n", user)
	user2, err := userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "2",
		Age:  38,
		Name: "Teste 2",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Segundo usuário criado: %v\n", user2)

	userResponse, err := userClient.GetUser(context.Background(), &pb.GetUserRequest{Id: "2"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Usuário obtido: %v\n", userResponse)
}
