package proto

import (
	"io"
	"log"
)

type HelloService struct {}

func (this *HelloService) HelloClientStream(stream HelloService_HelloClientStreamServer) error {
	for {
		req,err := stream.Recv()
		if err==nil {
			log.Println("Recv : ",req.Msg)
		}
		if err==io.EOF {
			continue
		}
	}
	return nil
}
