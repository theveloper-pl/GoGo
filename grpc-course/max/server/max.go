package main

import (
	"io"
	"log"

	pb "github.com/theveloper-pl/grpc-course/max/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"fmt"
)
func (s *Server) Max(stream pb.MaxService_MaxServer) error {
		res := int64(0)
		for {
			req, err := stream.Recv()
	
	
			if err == io.EOF {
				return nil
			}


			if err != nil {
				log.Fatalf("Error while reading the stream: %v\n", err)
			}


			log.Printf("Current res: %v, request: %v", res, req.Number)

			// Error handling

			if req.Number == 20 {
				return status.Errorf(
					codes.InvalidArgument,
					fmt.Sprintf("Error: Recived %v\n", req.Number),
				)
			}

			// Error handling

			if req.Number > res {
				res = req.Number
				err = stream.Send(&pb.MaxResponse{
					Number: res,
				})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}				
			}




		}


 return nil
}


