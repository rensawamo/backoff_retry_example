package main

import (
	"context"
	"log"
	"time"

	pb "backoff_retry_example/gen/buf/proto"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func retryableUnaryCall(client pb.SomeServiceClient, req *pb.SomeRequest) (*pb.SomeResponse, error) {
	var resp *pb.SomeResponse
	var err error

	b := backoff.NewExponentialBackOff()
	b.InitialInterval = 100 * time.Millisecond
	b.MaxInterval = 2 * time.Second
	b.MaxElapsedTime = 10 * time.Second

	operation := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		resp, err = client.SomeMethod(ctx, req)
		if err != nil {
			st, ok := status.FromError(err)
			if ok && isRetryableCode(st.Code()) {
				log.Printf("Retryable error: %v. Retrying...", err)
				return err
			}
			return backoff.Permanent(err)
		}
		return nil
	}

	err = backoff.Retry(operation, b)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func isRetryableCode(code codes.Code) bool {
	switch code {
	case codes.Unavailable, codes.DeadlineExceeded:
		return true
	default:
		return false
	}
}

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewSomeServiceClient(conn)

	req := &pb.SomeRequest{Message: "Hello, gRPC with Backoff!"}

	resp, err := retryableUnaryCall(client, req)
	if err != nil {
		log.Fatalf("Request failed after retries: %v", err)
	}

	log.Printf("Response received: %v", resp)
}
