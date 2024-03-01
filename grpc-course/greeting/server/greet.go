package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/theveloper-pl/grpc-course/greeting/proto"
)
func (s *Server) Greet(stream pb.GreetService_GreetServer) error {

		res := ""
		for {
			req, err := stream.Recv()
	
			if err == io.EOF {
				return stream.SendAndClose(&pb.GreetResponse{
					Message: res,
				})
			}
	
			if err != nil {
				log.Fatalf("Error while reading the stream: %v\n", err)
			}
			res = res + fmt.Sprintf("Hello %s!\n", req.Message)

		}


 return nil
}


