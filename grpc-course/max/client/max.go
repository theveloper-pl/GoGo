package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/theveloper-pl/grpc-course/max/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doMax(c pb.MaxServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Max: %v\n", err)
	}


	reqs := []*pb.MaxRequest {
		{Number: int64(1)},
		{Number: int64(5)},
		{Number: int64(3)},
		{Number: int64(6)},
		{Number: int64(2)},
		{Number: int64(20)},
		{Number: int64(22)},
		{Number: int64(4)},

	}

		waitc := make(chan struct{})
		go func() {
			for _, req := range reqs {
				log.Printf("Sending request: %v\n",req)
				stream.Send(req)
				time.Sleep(1*time.Second)
			}
			stream.CloseSend()
		}()

		go func(){
			for {
				res, err := stream.Recv()	
				
				if err == io.EOF {
					break
				}

				if err != nil {
					// log.Printf("Error while receiving %v\n", err)
					e, ok := status.FromError(err)

					if ok {
						log.Printf("Error message from server: %v\n", e.Message())
						log.Printf("Error code from server: %v\n", e.Code())

						if e.Code() == codes.InvalidArgument{
							fmt.Printf("Invalid argument")
						}

					}else {
						log.Fatalf("Not gRPC error %v", err)
					}

					break
				}

				log.Printf("Recived: %s\n", res)

			}
			close(waitc)
		}()

		<-waitc

}
