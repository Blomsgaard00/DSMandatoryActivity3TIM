

//publish method, 128 max characters
package main

import (
	proto "DSMandatoryActivity3TIM/gRPC"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client not working")
	}

	client := proto.NewChittyChatClient(conn)

	students, err := client.GetChatMessages(context.Background(), &proto.Empty{})
	if err != nil {
		log.Fatalf("Not working")
	}

	
}