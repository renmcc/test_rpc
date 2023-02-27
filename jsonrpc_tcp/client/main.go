package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"test-grpc/jsonrpc_tcp/service"
)

var _ service.HelloService = new(HelloServiceClient)

type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(request string, reply *string) error {
	err := c.client.Call(fmt.Sprintf("%s.Hello", service.HelloServiceName), "alice", &reply)
	if err != nil {
		return err
	}
	return nil
}

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	// 使用jsonrpc进行解码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{
		client: client,
	}, nil
}

func main() {
	client, err := NewHelloServiceClient("tcp", "192.168.10.1:1234")
	if err != nil {
		log.Fatal(err)
	}

	var reply string
	err = client.Hello("alice", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
