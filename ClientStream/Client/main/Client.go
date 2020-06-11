package main

import (
	services "Client/services"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn,err := grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	if err!=nil {
		log.Println(err)
	}
	defer conn.Close()
	c := services.NewHelloServiceClient(conn)
	stream,_ := c.HelloServiceStream(context.Background())
	for i:=0;i<10;i++ {
		stream.Send(&services.Request{Msg: "hello from client"})
		time.Sleep(time.Second)
	}
	res,err:=stream.CloseAndRecv()
	if err!=nil {
		log.Println(err)
	}
	log.Println("Recv : ",res.Msg)
}
