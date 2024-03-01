package main

import (
"log"
"google.golang.org/grpc"
"google.golang.org/grpc/credentials"
pb "github.com/theveloper-pl/grpc-course/max/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "D:/MyCode/golangv2/grpc-course/max/ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while using TLS: %v\n", err)
		}


		opts = append(opts, grpc.WithTransportCredentials(creds))

	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}

	defer conn.Close()

	c1 := pb.NewMaxServiceClient(conn)
	doMax(c1)


}