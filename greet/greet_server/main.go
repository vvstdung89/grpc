package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vvstdung89/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Printf("invoke with %v", req)
	firstname := req.GetGreeting().GetFirstName()
	result := "Hello " + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res,nil
}

func main() {
	fmt.Println("Hello")
	lis, err := net.Listen("tcp", "0.0.0.0:10001")
	if err != nil {
		log.Fatalf("Cannot listen on port 10001 %v", err)
	}
	
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
