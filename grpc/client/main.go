package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"test-grpc/grpc/middleware/auther"
	"test-grpc/grpc/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// grpc.Dial负责和gRPC服务建立链接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(auther.NewClientAuthentication("admin", "123456")))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// NewHelloServiceClient函数基于已经建立的链接构造HelloServiceClient对象,
	// 返回的client其实是一个HelloServiceClient接口对象
	client := pb.NewHelloServiceClient(conn)

	// 添加认证信息
	// crendential := metadata.MD{auther.ClientHeaderKey: []string{"admin"}, auther.ClientSecretKey: []string{"123456"}}
	// ctx := metadata.NewOutgoingContext(context.Background(), crendential)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs())

	// 通过接口定义的方法就可以调用服务端对应的gRPC服务提供的方法
	// 每次带上凭证信息
	reply, err := client.Hello(ctx, &pb.Request{Value: "张三"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

	// 客户端需要先调用Channel方法获取返回的流对象
	stream, err := client.Channel(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 在客户端我们将发送和接收操作放到两个独立的Goroutine。
	// 首先是向服务端发送数据
	go func() {
		for {
			if err := stream.Send(&pb.Request{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	// 然后在循环中接收服务端返回的数据
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
