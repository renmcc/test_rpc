package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"test-grpc/rpc_interface/service"
)

var _ service.HelloService = new(HelloService)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = fmt.Sprintf("hello, %s", request)
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	rpc.RegisterName("HelloService", new(HelloService))
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeConn(conn)
	}
}
