package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-study/grpc/helloworld/helloworld"
	"google.golang.org/grpc"
	"os"
	"time"
)

const address = "localhost:12347"
const defaultName = "test"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet %v", err)
	}
	log.Printf("Greeting %s", r.GetMessage())
}
