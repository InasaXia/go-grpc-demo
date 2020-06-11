package proto

import (
	"io"
	"log"
	"strconv"
)

type HelloService struct {}

func (this *HelloService) HelloServiceStream(stream HelloService_HelloServiceStreamServer) error {
	cnt := 0
	for {
		req,err := stream.Recv()
		if err==io.EOF {
			return stream.SendAndClose(&Response{Msg: strconv.Itoa(cnt)})
		}
		cnt++
		log.Println("Recv : ",req.Msg)
	}
	return nil
}
