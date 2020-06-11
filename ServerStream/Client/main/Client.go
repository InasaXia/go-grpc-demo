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
	stream,err := c.HelloServiceStream(context.Background(),&services.Request{Msg: "hello"})
	if err!=nil {
		log.Println(err)
	}
	for {
		res,err := stream.Recv()
		if err==io.EOF {
			log.Println("Server Response Done ...")
			break
		}
		log.Println("Recv : ",res.Msg)
	}
}
