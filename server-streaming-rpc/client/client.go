package client

import (
	"context"
	"fmt"
	"io"
	"server-streaming-rpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewStatusServiceClient(conn)

	stream, err := client.StreamStatus(context.Background(), &pb.StreamRequest{TaskId: "123"})
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			//servidor terminou de enviar o que precisava
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Printf("Recebendo status: %s, progresso %d%%\n", res.Message, res.Progress)
	}
}
