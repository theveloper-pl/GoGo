package main

import (
"log"
"google.golang.org/grpc"
"google.golang.org/grpc/credentials/insecure"
pb "github.com/theveloper-pl/grpc-course/greeting/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}

	defer conn.Close()

	c1 := pb.NewGreetServiceClient(conn)
	doGreet(c1)


}