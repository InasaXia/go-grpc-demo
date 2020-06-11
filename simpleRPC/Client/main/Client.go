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
	c := services.NewSimpleRPCServiceClient(conn)
	for i:=0;i<10;i++ {
		res,err := c.SimpleRPC(context.Background(),&services.Request{Msg: "hello"})
		if err!=nil {
			log.Println(err)
		}
		log.Println(res)
		time.Sleep(time.Second)
	}


}
