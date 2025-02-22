package main

import (
	// pb is the alias for the protocol buffer package generated from the service definition
	// grpc is the package used to set up the gRPC server

	"io"
	"log"

	pb "github.com/ShubhKanodia/goCPRg.git/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	//now will recieve list of helloRequest names from client and stream the response

	// var messages []string

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break //stop server sending the stream of responses as stream req from client is over
		}

		if err != nil {
			return err
		}

		log.Printf("Got req with name:%v", req.Name)

		resp := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}

		if err := stream.Send(resp); err != nil {
			return err
		}

	}

	return nil
}
