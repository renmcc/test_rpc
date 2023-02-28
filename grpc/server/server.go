package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"test-grpc/grpc/pb"

	"google.golang.org/grpc"
)

var _ pb.HelloServiceServer = new(HelloServiceServer)

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type HelloServiceServer struct {
	*pb.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Value: fmt.Sprintf("Hello %s", req.Value)}, nil
}

func main() {
	// 首先是通过grpc.NewServer()构造一个gRPC服务对象
	grpcServer := grpc.NewServer()
	// 然后通过gRPC插件生成的RegisterHelloServiceServer函数注册我们实现的HelloServiceImpl服务
	pb.RegisterHelloServiceServer(grpcServer, new(HelloServiceServer))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// 然后通过grpcServer.Serve(lis)在一个监听端口上提供gRPC服务
	grpcServer.Serve(lis)
}
