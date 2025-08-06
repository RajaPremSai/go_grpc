package main

import (
	"io"
	"log"

	pb "github.com/rajaprmesai/go_grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStraming(stream pb.GreetService_SayHelloBidirectionalStramingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return nil
		}
		log.Printf("Got request with name %v", req.Message)

		res := &pb.HelloResponse{
			Message: "Hello" + req.Message,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
