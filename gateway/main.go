package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/tattoo1880/my-grpc/gen"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// HTTP mux（gateway核心）
	mux := runtime.NewServeMux()

	// 连接 gRPC server
	err := pb.RegisterUserServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:50051",
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)
	if err != nil {
		log.Fatal("gateway register error:", err)
	}

	log.Println("HTTP gateway started at :8080")

	// 启动 HTTP server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
