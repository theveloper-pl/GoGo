package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/theveloper-pl/grpc-course/hello/proto"
)
func (s *Server) Hello(stream pb.HelloService_HelloServer) error {

		for {
			req, err := stream.Recv()
	
	
			if err == io.EOF {
				return nil
			}


			if err != nil {
				log.Fatalf("Error while reading the stream: %v\n", err)
			}
			res := fmt.Sprintf("Hello %s\n", req.FirstName)

			err = stream.Send(&pb.HelloResponse{
				Result: res,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}

		}


 return nil
}


