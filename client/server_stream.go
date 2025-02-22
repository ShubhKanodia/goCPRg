package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ShubhKanodia/goCPRg.git/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming has started!")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	//process the stream from the server on the client side

	for {
		//receive from the stream
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while streaming:  %v", err)
		}
		log.Println(message)
	}

	log.Printf("Streaming finished!")
}
