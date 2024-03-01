package main

import (
	"context"
	"log"

	pb "github.com/theveloper-pl/grpc-course/notify/proto"
)


func doNotify(c pb.NotifyServiceClient, my_text string) {
	log.Println("doNotify was invoked")
	res, err := c.Notify(context.Background(), &pb.NotifyRequest{
		Id:2137,
		Message: my_text,
	})

	if err != nil {
		log.Fatalf("Could not notify: %v\n", err)
	}

	log.Printf("Notification %s\n", res.Result)
}

