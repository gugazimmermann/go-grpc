package main

import (
	"context"
	"fmt"
	"go-grpc/greetpb/greetpb"
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

	doUnary(c)
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
		log.Fatalf("Error calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v\n", res.Result)
}
