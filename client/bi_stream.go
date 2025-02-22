package main

import (
	// pb is the alias for the protocol buffer package generated from the service definition
	// grpc is the package used to set up the gRPC server

	"context"
	"io"
	"log"
	"time"

	pb "github.com/ShubhKanodia/goCPRg.git/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming has started!")

	//create stream
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	//Cannot use for loop, that will make client and server responses coupled, making it similar to multiple unary req response cycles

	//hence will use go routines and channels
	waitc := make(chan struct{})

	//go routine to receive stream from server asynchronously
	go func() {
		for {
			message, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while streaming: %v", err)
			}
			log.Println(message)
		}

		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}

		time.Sleep(time.Second * 2) //simulate delay
	}
	//Signals to the server that the client has finished sending messages
	stream.CloseSend()

	//Blocks the main thread until the receiving goroutine closes the channel
	//so that all server responses are processed before the function returns
	//Prevents the program from terminating while still receiving messages

	<-waitc //everything has to come from server even after client has finished

	log.Printf("Bidirectional Streaming has finished!")

}
