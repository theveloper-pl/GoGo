package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "github.com/theveloper-pl/grpc-course/max/proto"

)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.MaxServiceServer
	
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listning on: %s\n", addr)

	tls := true
	opts := []grpc.ServerOption{}
	if tls {
		certFile := "D:/MyCode/golangv2/grpc-course/max/ssl/server.crt"
		keyFile := "D:/MyCode/golangv2/grpc-course/max/ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterMaxServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}