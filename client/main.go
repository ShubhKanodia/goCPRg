package main

import (
	"log"

	pb "github.com/ShubhKanodia/goCPRg.git/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	// Dial a gRPC server running on localhost at the specified port with insecure credentials (no encryption).
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close() //clost the connection at end of current function

	client := pb.NewGreetServiceClient(conn)

	//to stream hello to these names from the server
	names := &pb.NamesList{
		Names: []string{"John", "Mark", "Alice", "Bob"},
	}

	// callSayHello(client) //unary api call
	// callSayHelloServerStream(client, names) //server streaming call

	// callSayHelloClientStream(client, names) //client streaming call
	callSayHelloBidirectionalStream(client, names) //bidirectional -> both client server streaming call

}
