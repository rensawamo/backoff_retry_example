package main

import (
	"context"
	"log"
	"time"

	
	"google.golang.org/grpc"
)

func main() {
	// gRPCサーバへの接続設定
	conn, err := grpc.Dial("core-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewCoreServiceClient(conn)

	// リクエスト
	req := &pb.SomeRequest{Data: "example"}

	// 汎用リトライメソッドの呼び出し
	ctx := context.Background()
	resp, err := retryableCall(ctx, func(ctx context.Context) (*pb.SomeResponse, error) {
		return client.SomeMethod(ctx, req) // 呼び出したいgRPCメソッドを指定
	}, 7*time.Second, 100*time.Millisecond, 2*time.Second)

	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	log.Printf("Response received: %v", resp)
}
