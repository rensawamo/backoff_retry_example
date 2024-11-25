package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pb "backoff_retry_example/gen/buf/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SomeServiceServer is the server implementing the SomeService service.
type SomeServiceServer struct {
	pb.UnimplementedSomeServiceServer
}

// SomeMethod handles the incoming request and sometimes returns a retryable error.
func (s *SomeServiceServer) SomeMethod(ctx context.Context, req *pb.SomeRequest) (*pb.SomeResponse, error) {
	log.Printf("Received request: %v", req.Message)

	// Simulate random failure to trigger retries
	if rand.Intn(2) == 0 {
		log.Println("Simulating transient error: Unavailable")
		return nil, status.Error(codes.Unavailable, "simulated transient error")
	}

	// Simulate a delay to test DeadlineExceeded
	delay := time.Duration(rand.Intn(3)) * time.Second
	log.Printf("Simulating delay: %v", delay)
	time.Sleep(delay)

	// If the delay exceeds the client's timeout, gRPC automatically returns DeadlineExceeded.
	select {
	case <-ctx.Done():
		log.Println("Context deadline exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "operation timed out")
	default:
		// Success response
		resp := &pb.SomeResponse{
			Reply: "Hello from gRPC server!",
		}
		log.Printf("Sending response: %v", resp.Reply)
		return resp, nil
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Start listening on a port
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register the SomeService server
	pb.RegisterSomeServiceServer(grpcServer, &SomeServiceServer{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
