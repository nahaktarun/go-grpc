package hello

import (
	"context"
	"fmt"

	"github.com/nahaktarun/grpc-module2/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	proto.UnimplementedHelloServiceServer
}

func (s Service) SayHello(ctx context.Context, request *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {

	if request.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "Name cannot be empty")
	}
	return &proto.SayHelloResponse{Message: fmt.Sprintf("Hello %s", request.GetName())}, nil
}
