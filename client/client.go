package main

import (
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

	fmt.Println(c)

}
