package main

import (
	"context"
	"log"
	"time"

	pb "github.com/theveloper-pl/grpc-course/greeting/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	reqs := []*pb.GreetRequest{
		{Message: "Adam"},
		{Message: "Mateusz"},
		{Message: "Seba"},
	}

	stream, err := c.Greet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Greet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while closing Greet: %v\n", err)
	}

	log.Printf("Greet: %s\n", res.Message)

}
