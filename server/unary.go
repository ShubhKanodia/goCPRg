package main

import (
	"context"

	pb "github.com/ShubhKanodia/goCPRg.git/proto"
)

// unary server function
// method on the helloServer type, as helloServer implements the GreetService interface

// The context passed into the method allows the server to manage cancellation,
//
//	 timeouts, and deadlines for the request.
//
//	This is a common pattern in Go when working with networked operations.

// s *helloServer ptr cuz thre method is modifying the state of the server, so always use pointers throughout
// for overall consistency and correctness

// also for *pb.HelloResponse -> can pass by value but ptr more memory efficient in case of large structs
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {

	//This line returns a HelloResponse message with the field Message set to "Hello".
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
