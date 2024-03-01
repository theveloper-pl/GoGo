package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/theveloper-pl/grpc-course/hello/proto"
)

func doHello(c pb.HelloServiceClient) {
	log.Println("doHello was invoked")

	stream, err := c.Hello(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Hello: %v\n", err)
	}


	reqs := []*pb.HelloRequest {
		{FirstName: "Mateusz"},
		{FirstName: "Sbea"},
		{FirstName: "Adam"},
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
					log.Printf("Error while receiving %v\n", err)
					break
				}

				log.Printf("Recived: %s\n", res)

			}
			close(waitc)
		}()

		<-waitc

}
