package main

import (
	"context"
	proto "everywere/DSMandatoryActivity3TIM/gRPC"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ChittyChatServer struct {
	proto.UnimplementedChittyChatServer
	chatMessages []*proto.ChatMessage
}

func (s *ChittyChatServer) GetChatMessages(ctx context.Context, in *proto.Empty) (*proto.ChatMessages, error) {
	return &proto.ChatMessages{ChatMessages: s.chatMessages}, nil
}
func (s *ChittyChatServer) PostMessage(ctx context.Context, in *proto.ChatMessage) (*proto.Respons, error) {
	
	
	server.chatMessages = append(server.chatMessages,  )
	
	return , nil
}

func main() {
	server := &ChittyChatServer{chatMessages: []*proto.ChatMessage{}}

	server.start_server()

}

func (s *ChittyChatServer) start_server() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Server Error")
	}

	proto.RegisterChittyChatServer(grpcServer, s)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Server Error")
	}

}
