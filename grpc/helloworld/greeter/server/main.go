package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-study/grpc/helloworld/greeter/exception"
	"go-study/grpc/helloworld/helloworld"
	"google.golang.org/grpc"
	"net"
)

const port = ":12347"

// server is used to implement helloworld.GreeterServer.
type server struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("收到了 %v", in.GetName())

	return &helloworld.HelloReply{Message: "Hello  " + in.GetName()}, nil

}
func main() {
	lis, err := net.Listen("tcp", port)
	exception.Report(err, "failed to listen")
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
