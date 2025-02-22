# gRPC with Protocol Buffers in Go

## Overview
This project demonstrates the implementation of gRPC (Google Remote Procedure Call) using Protocol Buffers in Go, showcasing different types of API communication patterns: Unary, Server-side Streaming, Client-side Streaming, and Bidirectional Streaming.

## Why Protocol Buffers?
Protocol Buffers (protobufs) are Google's language-agnostic, platform-neutral extensible mechanism for serializing structured data. They offer several advantages over traditional data formats:

- **Performance**: Faster than JSON and XML due to binary serialization
- **Type Safety**: Strong typing prevents runtime errors
- **Schema Evolution**: Backward and forward compatibility
- **Code Generation**: Automatic generation of client and server code
- **Language Agnostic**: Support for multiple programming languages
- **Smaller Payload Size**: Efficient binary format reduces network bandwidth

## Communication Patterns Implemented

### 1. Unary API
- Traditional request-response pattern
- Client sends a single request, server responds with a single response
- Example: Simple greeting service

### 2. Server-side Streaming
- Client sends one request, server responds with a stream of messages
- Useful for scenarios like:
  - Real-time data updates
  - Progress updates for long-running operations
  - Continuous data feeds

### 3. Client-side Streaming
- Client sends a stream of messages, server responds with a single response
- Ideal for:
  - File uploads
  - Batch processing
  - Continuous data collection

### 4. Bidirectional Streaming
- Both client and server can send streams of messages independently
- Perfect for:
  - Real-time chat applications
  - Gaming
  - Interactive communication
  - Concurrent processing

## Implementation Details

### Protocol Buffer Definition
```protobuf
service GreetService {
    rpc SayHello(NoParam) returns (HelloResponse);
    rpc SayHelloServerStreaming(NamesList) returns (stream HelloResponse);
    rpc SayHelloClientStreaming(stream HelloRequest) returns (MessagesList);
    rpc SayHelloBidirectionalStreaming(stream HelloRequest) returns (stream HelloResponse);
}
```

### Key Features
- **Concurrent Processing**: Uses Go routines for handling streams
- **Channel-based Synchronization**: Ensures proper stream completion
- **Error Handling**: Robust error handling for stream operations
- **Simulated Delays**: Demonstrates realistic async communication

## Getting Started

### Prerequisites
- Go 1.16 or higher
- Protocol Buffers compiler (protoc)
- Go plugins for Protocol Buffers

### Installation
1. Clone the repository:
```bash
git clone https://github.com/ShubhKanodia/goCPRg.git
```

2. Install dependencies:
```bash
go mod tidy
```

3. Generate Protocol Buffer code:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/greet.proto
```

### Running the Application
1. Start the server:
```bash
cd server
go run *.go
```

2. In a new terminal, run the client:
```bash
cd client
go run *.go
```

## Best Practices Demonstrated
1. **Proper Stream Management**
   - Correct handling of stream closures
   - Error handling for stream operations
   - Use of channels for synchronization

2. **Code Organization**
   - Clear separation of concerns
   - Modular code structure
   - Clean protocol buffer definitions

3. **Resource Management**
   - Proper connection handling
   - Stream cleanup
   - Goroutine management
