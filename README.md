# Go gRPC Example Project

This project demonstrates how to implement and use gRPC services in Go, including unary, server-side streaming, client-side streaming, and bidirectional streaming RPCs.

## Project Structure

```
root/
├── client/         # gRPC client implementations
├── proto/          # Protocol Buffers definitions and generated code
├── server/         # gRPC server implementations
└── go.mod
```

## Getting Started

### Prerequisites

- Go 1.20+ installed
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

### Install Dependencies

```sh
go mod tidy
```

### Generate Go Code from Proto

If you modify `proto/greet.proto`, regenerate the Go code:

```sh
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

### Running the Server

Start the gRPC server:

```sh
cd server
go run .
```

### Running the Client

In a separate terminal, run the client:

```sh
cd client
go run .
```

You can uncomment the desired function call in `client/main.go` to test different RPC types:
- Unary: `callSayHello(client)`
- Server streaming: `callSayHelloServerStream(client, names)`
- Client streaming: `callSayHelloClientStream(client, names)`
- Bidirectional streaming: `callHelloBidirectionalStream(client, names)`

## gRPC Methods

Defined in [`proto/greet.proto`](proto/greet.proto):

- **SayHello**: Unary RPC
- **SayHelloServerSideStreaming**: Server-side streaming RPC
- **SayHelloClientSideStreaming**: Client-side streaming RPC
- **SayHelloBidirectionalStraming**: Bidirectional streaming RPC


## New Things Learnt

- **gRPC Streaming Types:**  
  Explored all four types of gRPC methods: unary, server-side streaming, client-side streaming, and bidirectional streaming.

- **Proto File Structure:**  
  Learned how to define services and messages in a `.proto` file and how field names/types map to generated Go code.

- **Code Generation:**  
  Understood the importance of regenerating Go code after modifying proto files using `protoc` with the appropriate plugins.

- **Method Naming:**  
  Realized that the method names in the Go server implementation must exactly match those generated from the proto file, or the server will panic at runtime.

- **Error Handling in Streams:**  
  Learned how to handle `io.EOF` and other errors properly in streaming RPCs, both on the client and server side.

- **Appending to Slices:**  
  Noted the difference between `append(messages, "Hello", req.Message)` (which appends two elements) and `append(messages, "Hello "+req.Message)` (which appends one combined string).

- **Debugging gRPC Panics:**  
  Understood how mismatches between proto definitions and Go implementations can cause nil pointer dereferences and how to debug them.


## Learning Resources

- [gRPC in Go](https://grpc.io/docs/languages/go/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)

---