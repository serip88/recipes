package main

import (
	"context"
	"net"

	proto "github.com/serip88/recipes/fiber-grpc-rz/protogen/service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedAddServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *server) Add(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}
