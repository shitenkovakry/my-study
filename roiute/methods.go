package roiute

import (
	"context"
	"my-study/mygrpc"
)

type Server struct {
	mygrpc.UnimplementedRoiuteMessageServer
}

func New() *Server {
	return &Server{}
}

func (server *Server) GetMessage(ctx context.Context, request *mygrpc.HelloRequest) (*mygrpc.HelloResponse, error) {
	response := &mygrpc.HelloResponse{
		Message: "Hello, " + request.Name,
	}

	return response, nil
}
