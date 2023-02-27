package service

const (
	HelloServiceName = "HelloService"
)

type HelloService interface {
	Hello(request string, reply *string) error
	Calc(req *CalcRequest, reply *int) error
}

type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}
