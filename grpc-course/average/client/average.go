package main

import (
	"context"
	"log"
	"time"

	pb "github.com/theveloper-pl/grpc-course/average/proto"
)

func doAverage(c pb.AverageServiceClient) {
	log.Println("doAverage was invoked")

	reqs := []*pb.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
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

	log.Printf("Average: %s\n", res.Number)

}
