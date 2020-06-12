package main

import (
	services "Client02/services"
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
	c := services.NewHelloServiceClient(conn)
	stream,err := c.HelloClientStream(context.Background())
	if err!=nil {
		log.Println(err)
	}
	for {
		stream.Send(&services.Request{Msg: "hello from Client"})
		time.Sleep(time.Second)
	}
}