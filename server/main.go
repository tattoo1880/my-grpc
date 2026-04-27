package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

package main

import (
"context"
"log"
"net"

pb "github.com/tattoo1880/my-grpc/proto"

"google.golang.org/grpc"
)

// server 实现 gRPC 接口
type server struct {
	pb.UnimplementedUserServiceServer
}

// 实现 RPC 方法
func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	log.Println("收到请求 user id:", req.Id)

	return &pb.UserResponse{
		Id:   req.Id,
		Name: "Alice from gRPC server",
	}, nil
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	// 创建 gRPC server
	grpcServer := grpc.NewServer()

	// 注册服务
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("gRPC server started at :50051")

	// 启动
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}