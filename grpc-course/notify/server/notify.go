package main

import (
	"context"
	"log"

	pb "github.com/theveloper-pl/grpc-course/notify/proto"
)

func (s *Server) Notify(ctx context.Context, in *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	log.Printf("Notify function was invoked with %v\n", in)
	return &pb.NotifyResponse {
		Result: "Hello " + in.Message,
	}, nil
}


