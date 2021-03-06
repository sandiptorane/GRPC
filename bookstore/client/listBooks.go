package main

import (
	books "bookstore/bookproto"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main(){
	var conn *grpc.ClientConn
	//dial for server conn
	conn,err:= grpc.Dial(":8082",grpc.WithInsecure())
	if err!=nil{
		log.Fatal("did not connect : ",err)
	}
	defer conn.Close()

	//initialize bookService
	bookClient := books.NewBookServiceClient(conn)

	response,err :=bookClient.ListBooks(context.Background(),&books.ListBookRequest{})
	if err!=nil{
		fmt.Println(err)
		return
	}
	book := response.GetBooks()
	fmt.Println("books present in bookstore:\n")
	for i,_:=range book{
		fmt.Println(book[i])
	}

}
