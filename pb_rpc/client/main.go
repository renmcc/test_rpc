package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"test-grpc/pb_rpc/codec"
	"test-grpc/pb_rpc/service"
)

var _ service.HelloService = new(HelloServiceClient)

type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(request *service.Request, response *service.Response) error {
	err := c.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)
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
	// 使用protobuf进行解码
	client := rpc.NewClientWithCodec(codec.NewClientCodec(conn))
	return &HelloServiceClient{
		client: client,
	}, nil
}

func main() {
	client, err := NewHelloServiceClient("tcp", "192.168.10.1:1234")
	if err != nil {
		log.Fatal(err)
	}

	var reply = new(service.Response)
	err = client.Hello(&service.Request{Value: "bob1234"}, reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.Value)
}
