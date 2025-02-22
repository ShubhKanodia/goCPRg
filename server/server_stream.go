package main

import (
	"log"
	"time"

	// pb is the alias for the protocol buffer package generated from the service definition
	// grpc is the package used to set up the gRPC server
	pb "github.com/ShubhKanodia/goCPRg.git/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with   names : %v", req.Names)

	for _, name := range req.Names {
		resp := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		//use stream.Send as we are sending response as a stream to the client
		if err := stream.Send(resp); err != nil {
			return err
		}

		time.Sleep(2 * time.Second) //to simulate the delay
	}
	return nil
}
