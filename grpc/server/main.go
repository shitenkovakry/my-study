package main

import "my-study/mygrpc"

//import "context"

type Server struct{}

//func (server *Server) SayHello(ctx context.Context, request )

func main() {
	mygrpc.RegisterRoiuteMessageServer()
}
