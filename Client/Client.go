// publish method, 128 max characters
package main

import (
	proto "DSMandatoryActivity3TIM/gRPC"
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var timestamp int32


func main() {
	conn, err := grpc.NewClient("localhost:5100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client not working")
	}
	client := proto.NewChittyChatClient(conn)
	
	timestamp = 1
	connectionMessage := &proto.Connect{
		Active : true,
		Timestamp: timestamp,
	}

	stream, err := client.CreateStream(context.Background(), connectionMessage )
	if err != nil {
		log.Fatalf("Not working")
	}

	
	

	for{
		random := rand.IntN(1000)
		input, err := stream.Recv()
		if err != nil {
			log.Fatalf("Not working")
		}
		fmt.Println(input.Message)
		
		if(input.Timestamp > timestamp){
			timestamp = input.Timestamp
		}
		sendMessage := &proto.Message{
			Id : 1,
			Message : "Random valid message",
			Timestamp : timestamp,
		}
		if(random == 1){
			stream.SendMsg(sendMessage)
		}
	}

	
}