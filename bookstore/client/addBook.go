package main

import (
	books "bookstore/bookproto"
	"context"
	"fmt"
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

	//initialize BookService
	bookClient := books.NewBookServiceClient(conn)
	book := &books.AddBookRequest{
		Book: &books.Book{
			Title:         "clean code",
			Author:        "uncle bob",
			Content:       "how to do clean code",
		},
	}
	response,err :=bookClient.AddBook(context.Background(),book)
	if err!=nil{
		log.Fatal("error when calling CreateBook:",err)
	}
	addedBook := response.GetBook()

	fmt.Println("New book added to bookstore:",addedBook)
}
