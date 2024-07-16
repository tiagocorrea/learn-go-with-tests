package main

import (
	"context"
	"log"
	"net"

	"github.com/tiagocorrea/go-specs-greet/adapters/grpcserver"
	"github.com/tiagocorrea/go-specs-greet/domain/interactions"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &GreeterServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type GreeterServer struct {
	grpcserver.UnimplementedGreeterServer
}

func (g GreeterServer) Greet(ctx context.Context, request *grpcserver.GreetRequest) (*grpcserver.GreetReply, error) {
	return &grpcserver.GreetReply{Message: interactions.Greet(request.Name)}, nil
}
