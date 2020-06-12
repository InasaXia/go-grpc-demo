package main

import (
	services "Client/services"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn,err := grpc.Dial("rhel:8081",grpc.WithInsecure())
	if err!=nil {
		log.Println(err)
	}
	defer conn.Close()
	c := services.NewHelloServiceClient(conn)
	stream,_ := c.HelloServiceStream(context.Background())
	for i:=0;i<1000;i++ {
		stream.Send(&services.Request{Msg: "hello from client"})
		time.Sleep(time.Second)
	}
	res,err:=stream.CloseAndRecv()
	if err!=nil {
		log.Println(err)
	}
	log.Println("Recv : ",res.Msg)
}
