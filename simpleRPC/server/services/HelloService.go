package proto

import (
	"context"
	"log"
)

type HelloService struct {}

func (this *HelloService) SimpleRPC(ctx context.Context,req *Request) (*Response, error) {
	log.Println("Recv : ",req.Msg)
	return &Response{Msg: req.Msg},nil
}
