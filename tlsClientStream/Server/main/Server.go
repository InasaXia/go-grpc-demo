package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net"
	services "Server/services"
)

var CA_PATH = "/opt/tls/tlsWithPhrase/ca.crt"
var SERVER_CERT = "/opt/tls/tlsWithPhrase/server.crt"
var SERVER_KEY = "/opt/tls/tlsWithPhrase/server.key"
func main() {
	pool := x509.NewCertPool()
	caCrt,err := ioutil.ReadFile(CA_PATH)
	if err!=nil {
		panic(err)
	}
	pool.AppendCertsFromPEM(caCrt)
	keyByte,err := ioutil.ReadFile(SERVER_KEY)
	certS,err := ioutil.ReadFile(SERVER_CERT)
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
		Certificates:       []tls.Certificate{certificate},
		ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          pool,
		InsecureSkipVerify: false,
	}
	listen,_ := net.Listen("tcp","rhel:8081")
	defer listen.Close()
	rpcServer := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))
	services.RegisterHelloServiceServer(rpcServer,new(services.HelloService))
	rpcServer.Serve(listen)
}
