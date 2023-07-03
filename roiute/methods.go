package roiute

import (
	"context"
	"fmt"
	mygrpc "my-study/mygrpc"
)

type Server struct {
	mygrpc.UnimplementedRoiuteMessageServer
}

func New() *Server {
	return &Server{}
}

func (server *Server) SayHello(ctx context.Context, request *mygrpc.HelloRequest) (*mygrpc.HelloResponse, error) {
	response := &mygrpc.HelloResponse{
		Message: fmt.Sprintf("Hello, %s. You said: %s", request.Name, request.Message),
	}

	return response, nil
}
