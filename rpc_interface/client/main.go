package main

import (
	"fmt"
	"log"
	"net/rpc"
	"test-grpc/rpc_interface/service"
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

func NewHelloServiceClient(net, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(net, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		client: client,
	}, nil
}

func main() {
	client, err := NewHelloServiceClient("tcp", "localhost:1234")
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
