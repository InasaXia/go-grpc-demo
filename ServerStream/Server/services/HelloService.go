package proto

import (
	"log"
	"time"
)

type HelloService struct {}

func (this *HelloService) HelloServiceStream(req *Request,stream HelloService_HelloServiceStreamServer) error {
	tmp := req.Msg
	log.Println("Recv : ",tmp)
	for i:=0;i<10;i++ {
		stream.Send(&Response{Msg: tmp+" from Server"})
		time.Sleep(time.Second)
	}
	return nil
}