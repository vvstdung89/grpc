package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vvstdung89/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello")
	lis, err := net.Listen("tcp", "0.0.0.0:10001")
	if err != nil {
		log.Fatalf("Cannot listen on port 10001 %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer()

}
