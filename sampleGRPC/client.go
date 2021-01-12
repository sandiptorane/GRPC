package main

import (
	"sampleGRPC/chat"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main(){
	var conn *grpc.ClientConn
    //dial server it will returns connection
	conn,err:= grpc.Dial(":8082",grpc.WithInsecure())
	if err!=nil{
		log.Fatal("did not connect : ",err)
	}
	defer conn.Close()

	//initialize chat Message
	ch := chat.NewChatServiceClient(conn)
	response,err :=ch.SayHello(context.Background(),&chat.Message{Body: "Hello from client"})
	if err!=nil{
		log.Fatal("error when calling SayHello:",err)
	}
	fmt.Println("Response from server:",response.Body)
}
