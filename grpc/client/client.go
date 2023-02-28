package main

import (
	"context"
	"fmt"
	"log"
	"test-grpc/grpc/pb"

	"google.golang.org/grpc"
)

func main() {
	// grpc.Dial负责和gRPC服务建立链接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// NewHelloServiceClient函数基于已经建立的链接构造HelloServiceClient对象,
	// 返回的client其实是一个HelloServiceClient接口对象
	client := pb.NewHelloServiceClient(conn)

	// 通过接口定义的方法就可以调用服务端对应的gRPC服务提供的方法
	req := &pb.Request{Value: "张三"}
	reply, err := client.Hello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
