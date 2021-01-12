package main

import (
	"sampleGRPC/chat"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("gRPC server running!")
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("Listen err on port 8082 :", err)
	}

    s := chat.Server{}

	//initialize new grpc server
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer,&s)
	if err:= grpcServer.Serve(listen); err!=nil{
		log.Fatal("failed to serve:",err)
	}
}
