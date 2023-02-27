package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"test-grpc/jsonrpc_http/service"
)

var _ service.HelloService = new(HelloService)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = fmt.Sprintf("hello, %s", request)
	return nil
}

func (s *HelloService) Calc(req *service.CalcRequest, reply *int) error {
	*reply = req.A + req.B
	return nil
}

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

func NewRPCReadWriteCloserFromHTTP(w http.ResponseWriter, r *http.Request) *RPCReadWriteCloser {
	return &RPCReadWriteCloser{w, r.Body}
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		// 使用jsonrpc进行编码
		rpc.ServeCodec(jsonrpc.NewServerCodec(NewRPCReadWriteCloserFromHTTP(w, r)))
	})

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}
