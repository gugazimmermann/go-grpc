package main

import (
	"context"
	"fmt"
	"go-grpc/greetpb/greetpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client test...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	// doUnary(c)

	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	log.Println("Starting Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Guga",
			LastName:  "Zimmermann",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Printf("Error calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v\n", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	log.Println("Starting Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Guga",
			LastName:  "Zimmermann",
		},
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Printf("Error calling Server Streaming GreetManyTimes RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v\n", res.GetResult())
	}
}
