package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"test-grpc/grpc/middleware/auther"
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

func (s *HelloServiceServer) Channel(stream pb.HelloService_ChannelServer) error {
	// 服务端在循环中接收客户端发来的数据
	for {
		// 接收一个请求
		args, err := stream.Recv()
		if err != nil {
			log.Println(err.Error())
			// 如果遇到io.EOF表示客户端流被关闭
			if err == io.EOF {
				return nil
			}
			return err
		}

		// 响应一个请求
		// 生成返回的数据通过流发送给客户端
		resp := &pb.Response{Value: "hello:" + args.GetValue()}
		err = stream.Send(resp)
		if err != nil {
			// 服务端发送异常, 函数退出, 服务端流关闭
			return err
		}
	}
}

func main() {
	// 首先是通过grpc.NewServer()构造一个gRPC服务对象
	grpcServer := grpc.NewServer(
		// 添加认证中间件, 如果有多个中间件需要添加 使用ChainUnaryInterceptor
		grpc.UnaryInterceptor(auther.GrpcAuthUnaryServerInterceptor()),
		// 添加stream API的拦截器
		grpc.StreamInterceptor(auther.GrpcAuthStreamServerInterceptor()),
	)
	// 然后通过gRPC插件生成的RegisterHelloServiceServer函数注册我们实现的HelloServiceImpl服务
	pb.RegisterHelloServiceServer(grpcServer, new(HelloServiceServer))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// 然后通过grpcServer.Serve(lis)在一个监听端口上提供gRPC服务
	grpcServer.Serve(lis)
}
