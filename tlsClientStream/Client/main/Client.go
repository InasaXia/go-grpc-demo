package main

import (
	services "Client/services"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	caCrtPath := "/opt/tls/tlsWithPhrase/ca.crt"
	clientCrt := "/opt/tls/tlsWithPhrase/client.crt"
	clientKey := "/opt/tls/tlsWithPhrase/client.key"
	pool := x509.NewCertPool()
	caCrt,err := ioutil.ReadFile(caCrtPath)
	if err!=nil {
		panic(err)
	}
	pool.AppendCertsFromPEM(caCrt)
	keyByte,err := ioutil.ReadFile(clientKey)
	certS,err := ioutil.ReadFile(clientCrt)
	keyBlock,_ := pem.Decode(keyByte)
	keyDER,err := x509.DecryptPEMBlock(keyBlock,[]byte("test"))
	if err!=nil {
		panic(err)
	}
	keyBlock.Bytes=keyDER
	keyBlock.Headers=nil
	keyPem := pem.EncodeToMemory(keyBlock)
	certificate,err := tls.X509KeyPair(certS,keyPem)
	if err!=nil {
		panic(err)
	}
	tlsConfig := &tls.Config{
		Certificates:                []tls.Certificate{certificate},
		RootCAs:                     pool,
		InsecureSkipVerify: 		 false,
		ServerName: 				 "rhel",
		MaxVersion: 				 tls.VersionTLS13,
		MinVersion: 				 tls.VersionTLS13,
	}
	conn,err := grpc.Dial("rhel:8081",grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	if err!=nil {
		log.Println(err)
	}
	defer conn.Close()
	c := services.NewHelloServiceClient(conn)
	stream,err := c.HelloClientStream(context.Background())
	for {
		stream.Send(&services.Request{Msg: "hello from Client"})
		time.Sleep(time.Second)
	}
}
