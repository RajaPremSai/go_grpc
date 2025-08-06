package main

import (
	"context"
	"io"
	"log"

	pb "github.com/rajaprmesai/go_grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming Started")

	stream, err := client.SayHelloServerSideStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("could not send names :%v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break

		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)

		}
		log.Println(message)
	}
	log.Printf("Streaming Finished")
}
