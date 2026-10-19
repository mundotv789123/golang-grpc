package server

import (
	"fmt"
	"net"
	"server-streaming-rpc/pb"
	"time"

	"google.golang.org/grpc"
)

func Run() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterStatusServiceServer(s, &StatusService{})
	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}

type StatusService struct {
	pb.UnimplementedStatusServiceServer
}

func (s *StatusService) StreamStatus(req *pb.StreamRequest, stream pb.StatusService_StreamStatusServer) error {
	taskId := req.TaskId

	fmt.Printf("Streaming status task id: %s\n", taskId)

	for i := 0; i < 100; i += 10 {

		status := pb.StreamResponse{
			Message:  "Progresso",
			Progress: int32(i),
		}

		err := stream.Send(&status)
		if err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}
