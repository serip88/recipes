package main

import (
	"context"
	"fmt"
	"net"

	"github.com/serip88/recipes/fiber-grpc-rn/database"

	"github.com/serip88/recipes/fiber-grpc-rn/server/handler"
	proto "github.com/serip88/recipes/protogen/service/v1"
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
	//conect db
	database.ConnectDB()
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *server) Add(_ context.Context, request *proto.Request) (*proto.Response, error) {
	println("B Add func...")
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}
func (s *server) GetUser(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	println("B GetUser func...")
	email := request.GetUser().GetEmail()
	println("email...", email)
	mUser, err := handler.GetUserByEmail(email)
	if err != nil {
		return nil, err
	} else if mUser == nil {
		return nil, nil
	}
	fmt.Println("mUser....", mUser)
	user := &proto.User{
		Id:       mUser.ID.String(),
		Email:    mUser.Email,
		Password: mUser.Password,
		Name:     mUser.Names,
	}
	println("E GetUser func...")
	return &proto.Response{User: user}, nil
}
