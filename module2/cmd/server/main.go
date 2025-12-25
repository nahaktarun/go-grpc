package main

import (
	"log"
	"net"

	"github.com/nahaktarun/grpc-module2/internal/todo"
	"github.com/nahaktarun/grpc-module2/proto"
	"google.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()

	// helloService := hello.Service{}
	todoService := todo.NewService()

	// proto.RegisterHelloServiceServer(grpcServer, helloService)

	proto.RegisterTodoServiceServer(grpcServer, todoService)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting server on the port: :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
