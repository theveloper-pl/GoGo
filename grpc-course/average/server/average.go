package main

import (
	"io"
	"log"

	pb "github.com/theveloper-pl/grpc-course/average/proto"
)
func (s *Server) Average(stream pb.AverageService_AverageServer) error {

		var sum float32 = 0
		var count float32 = 0
		for {
			req, err := stream.Recv()
	
			if err == io.EOF {
				avg := float32(sum / count)
				return stream.SendAndClose(&pb.AverageResponse{
					Number: avg,
				})
			}
	
			if err != nil {
				log.Fatalf("Error while reading the stream: %v\n", err)
			}
			sum = sum + float32(req.Number)
			count+=1

		}


 return nil
}


