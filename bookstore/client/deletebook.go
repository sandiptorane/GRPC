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
    //delete book
	response,err :=bookClient.DeleteBook(context.Background(),&books.DeleteBookRequest{Id: 2})
	if err!=nil{
		fmt.Println(err)
		return
	}
	success := response.GetSuccess()
	if success == true {
		fmt.Println("book deleted successfully")
		return
	}


}
