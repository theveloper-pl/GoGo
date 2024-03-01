package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/theveloper-pl/grpc-course/average/proto"

)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.AverageServiceServer
	
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listning on: %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterAverageServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}