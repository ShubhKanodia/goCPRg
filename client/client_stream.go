package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ShubhKanodia/goCPRg.git/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming has started!")

	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}

		log.Printf("Send the request with the name: %s", name)

		time.Sleep(2 * time.Second)
	}

	//u gotta close the stream and recieve the response
	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished!")

	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	log.Printf("%v", res.Message)
}
