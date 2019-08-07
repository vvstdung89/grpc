package main

import (
	"context"
	"fmt"
	"github.com/vvstdung89/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)


func main() {
	fmt.Println("Hello I am a client ")
	cc,err := grpc.Dial("localhost:10001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cc.Close()
	
	c := greetpb.NewGreetServiceClient(cc)
	
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dung",
			LastName: "Van",
		},
	}
	res,err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error", err)
	}
	log.Printf("Response %v", res.Result)
}
