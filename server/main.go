package main

import (
	"context"
	"net"

	service "github.com/skyapps-id/go-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

type person struct {
	name string
	age  int64
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	service.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *service.Request) (*service.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &service.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *service.Request) (*service.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &service.Response{Result: result}, nil
}

func (s *server) User(ctx context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	a := request.GetName()

	result := person{
		name: "none",
		age:  0,
	}

	if a == "aji" {
		result = person{
			name: "Aji",
			age:  22,
		}
	}

	return &service.UserResponse{Name: result.name, Age: result.age}, nil
}
