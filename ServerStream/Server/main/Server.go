package main

import (
	services "Server/services"
	"google.golang.org/grpc"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterHelloServiceServer(rpcServer,new(services.HelloService))
	listen,_ := net.Listen("tcp",":8080")
	rpcServer.Serve(listen)
}
