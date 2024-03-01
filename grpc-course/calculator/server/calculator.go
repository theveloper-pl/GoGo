package main

import (
	// "fmt"
	"log"

	pb "github.com/theveloper-pl/grpc-course/calculator/proto"
)

func (s *Server) CalculateManyTimes(in *pb.CalculatorRequest, stream pb.CalculatorService_CalculateManyTimesServer) error {
	log.Printf("Calculator stream was invoked with %v\n", in)

		k := int64(2);
		N := in.Number;
		for N > 1{
			if (N % k) == 0{
				stream.Send(&pb.CalculatorResponse{Result: k,})
			N = N / k 
			}else{
				k = k + 1
			}
		}


 return nil
}



