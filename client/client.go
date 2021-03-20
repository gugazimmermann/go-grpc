package main

import (
	"context"
	"fmt"
	"go-grpc/greetpb/greetpb"
	"io"
	"log"
	"time"

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

	// doServerStreaming(c)

	doClientStreaming(c)
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

func doClientStreaming(c greetpb.GreetServiceClient) {
	log.Println("Starting Client Streaming RPC...")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Printf("Error while calling LongGreet: %v", err)
	}

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Guga",
				LastName:  "Zimmermann",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "José",
				LastName:  "Augusto",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Renata",
				LastName:  "Negreiros",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Felipe",
				LastName:  "Kauling",
			},
		},
	}

	for _, req := range requests {
		log.Printf("Seding req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error while receiving response from LongGreet: %v", err)
	}

	log.Printf("LongGreet response: %v", res)
}
