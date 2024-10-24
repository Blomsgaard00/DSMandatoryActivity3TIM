package main

import (
	proto "DSMandatoryActivity3TIM/gRPC"
	"context"
	"fmt"
	"sync"
	"log"
	"net"
	"google.golang.org/grpc"
)
//Heavily inspired by https://medium.com/@bhadange.atharv/building-a-real-time-chat-application-with-grpc-and-go-aa226937ad3c
type Connection struct {
	proto.UnimplementedChittyChatServer
	stream proto.ChittyChat_CreateStreamServer
	name string
	active bool
	error  chan error
}

type Pool struct {
	proto.UnimplementedChittyChatServer
	Connection []*Connection
 }


func (p *Pool) CreateStream(pconn *proto.Connect, stream proto.ChittyChat_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		name: pconn.User.Name,
		active: true,
		error:  make(chan error),
	}
   
	p.Connection = append(p.Connection, conn)
	
	return <-conn.error
}

func (s *Pool) BroadcastMessage(ctx context.Context, msg *proto.Message) (*proto.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)
   
	for _, conn := range s.Connection {
	 wait.Add(1)
   
	 go func(msg *proto.Message, conn *Connection) {
	  defer wait.Done()
   
	  if conn.active {
	   err := conn.stream.Send(msg)
	   fmt.Printf("Sending message to: %v from %v", conn.name, msg.Id) // update this
   
	   if err != nil {
		fmt.Printf("Error with Stream: %v - Error: %v\n", conn.stream, err)
		conn.active = false
		conn.error <- err
	   }
	  }
	 }(msg, conn)
   
	}
   
	go func() {
	 wait.Wait()
	 close(done)
	}()
   
	<-done
	return &proto.Close{}, nil
   }


   func main() {
		// Create a new gRPC server
		grpcServer := grpc.NewServer()
	
		// Create a new connection pool
		var conn []*Connection
	
		pool := &Pool{
		Connection: conn,
		}
	
		// Register the pool with the gRPC server
		proto.RegisterChittyChatServer(grpcServer, pool)
	
		// Create a TCP listener at port 8080
		listener, err := net.Listen("tcp", ":8080")
	
		if err != nil {
		log.Fatalf("Error creating the server %v", err)
		}
	
		fmt.Println("Server started at port :8080")
	
		// Start serving requests at port 8080
		if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error creating the server %v", err)
		}
   }
