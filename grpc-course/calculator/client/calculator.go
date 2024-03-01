package main

import (
	"context"
	"io"
	"log"

	pb "github.com/theveloper-pl/grpc-course/calculator/proto"
)

func doCalculateManyTimes(c pb.CalculatorServiceClient) {
	log.Println("doCalculate was invoked")

	req := &pb.CalculatorRequest{
		Number: 12390392840,
	}

	stream, err := c.CalculateManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling CalculateManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("CalculateManyTimes: %v\n", msg.Result)
	}

}
