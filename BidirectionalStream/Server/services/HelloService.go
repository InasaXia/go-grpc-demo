package proto

import (
	"io"
	"log"
	"time"
)

type HelloService struct {}

func (this *HelloService) HelloBidirectionalStream(stream HelloService_HelloBidirectionalStreamServer) error {
	for {
		res,err := stream.Recv()
		if err==io.EOF {
			return nil
		}
		log.Println("Recv : ",res.Msg)
		stream.Send(&Response{Msg: "msg from Server"})
		time.Sleep(time.Second)
	}
	return nil
}
