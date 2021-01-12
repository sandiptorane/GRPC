package chat

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {

}

func (s *Server)SayHello(ctx context.Context, msg *Message)(*Message,error){
	log.Printf("Receive message body from client: %s", msg.Body)
	return &Message{Body: "Hello From the Server!"}, nil //Message is generated in chat.pb.go
}
