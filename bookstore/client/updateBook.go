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
	//dial server it will returns connection
	conn,err:= grpc.Dial(":8082",grpc.WithInsecure())
	if err!=nil{
		log.Fatal("did not connect : ",err)
	}
	defer conn.Close()

	//initialize BookService
	bookClient := books.NewBookServiceClient(conn)
	book := &books.UpdateBookRequest{
		Book: &books.Book{
			Id: 1,
			Title:         "clean code",
			Author:        "uncle bob",
			Content:       "how to do clean code",
		},
	}
	response,err :=bookClient.UpdateBook(context.Background(),book)
	if err!=nil{
		log.Fatal(err)
	}
	addedBook := response.GetBook()

	fmt.Println("book updated successfully on the bookstore:",addedBook)
}
