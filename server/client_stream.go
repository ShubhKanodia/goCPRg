package main

import (
	// pb is the alias for the protocol buffer package generated from the service definition
	// grpc is the package used to set up the gRPC server
	"io"
	"log"

	pb "github.com/ShubhKanodia/goCPRg.git/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		//how to process stream? -> u do stream.Recv()
		req, err := stream.Recv()
		if err == io.EOF {
			//send the answer from the server
			return stream.SendAndClose(&pb.MessagesList{
				Message: messages,
			})
		}
		if err != nil {
			return err
		}

		log.Printf("Got req with name:%v", req.Name)

		messages = append(messages, "Hello", req.Name)
	}
}
