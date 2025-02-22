package main

import (
	// Importing the context package allows the client to
	// control the duration of the gRPC call with a timeout or deadline.
	"context"
	"log"
	"time"

	pb "github.com/ShubhKanodia/goCPRg.git/proto"
)

// GreetServiceClient is an interface here
// interfaces are generally passed by value
// When you pass an interface by value,
// it behaves differently from regular structs because the Go runtime manages the interface internally.

// Go interfaces are reference types,
// and they internally hold a pointer to the actual implementation

func callSayHello(client pb.GreetServiceClient) {
	//context.Background() -> base context used to create new context
	//syntax is -> "context.WithTimeout(parent, timeout)"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	//cancel() is a function that is part of the context package,
	// and it’s used to cancel or abort an ongoing operation that was initiated with a context.
	// Ensure cancel is called to release resources when done
	// cancel is generally used with with context.WithTimeout() or context.WithCancel()
	defer cancel()

	// This line calls the SayHello method on the client.
	// The ctx ensures the call respects the timeout (1 second).
	// &pb.NoParam{} creates a new empty NoParam message to send as the request, as defined in the protocol file.
	resp, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", resp.Message) //print the response from server
}
