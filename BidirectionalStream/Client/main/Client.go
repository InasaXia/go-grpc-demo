package main

import (
	services "Client/services"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn,err := grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	if err!=nil {
		log.Println(err)
	}
	defer conn.Close()
	c := services.NewHelloServiceClient(conn)
	stream,_ := c.HelloBidirectionalStream(context.Background())
	for {
		stream.Send(&services.Request{Msg: "msg from Client"})
		res,err:=stream.Recv()
		if err==io.EOF {
			continue
		}
		log.Println("Recv : ",res.Msg)
	}
}
