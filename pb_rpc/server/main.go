package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"test-grpc/pb_rpc/codec"
	"test-grpc/pb_rpc/service"
)

var _ service.HelloService = new(HelloService)

type HelloService struct {
}

func (s *HelloService) Hello(request *service.Request, reply *service.Response) error {
	reply.Value = fmt.Sprintf("hello, %s", request.Value)
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	rpc.RegisterName(service.SERVICE_NAME, new(HelloService))
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 使用protobuf进行编码
		go rpc.ServeCodec(codec.NewServerCodec(conn))
	}
}
