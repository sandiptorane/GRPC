package main

import (
	books "bookstore/bookproto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("gRPC bookService server running!")
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("Listen err on port 8082 :", err)
	}
	s := Server{}

	s.initializeServer() //add some books to server

	//initialize new grpc server
	grpcServer := grpc.NewServer()
	books.RegisterBookServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("failed to serve:", err)
	}
}

func (s *Server) initializeServer() {
	book1 := &books.Book{
		Id:      1,
		Title:   "abc",
		Author:  "xyz",
		Content: "sggr",
	}
	book2 := &books.Book{
		Id:      2,
		Title:   "dfg",
		Author:  "me",
		Content: "something",
	}
	var bookStore []*books.Book
	bookStore = append(bookStore, book1, book2)
	s.bookStore = bookStore
}
