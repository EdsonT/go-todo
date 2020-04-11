package main

import (
	"fmt"
	"go-todo/todo"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis,err:=net.Listen("tcp",fmt.Sprintf(":%d",14586))
	if err != nil {
		log.Fatalf("Failed to listen: %v",err)
	}
	s:=todo.Server{}
	grpcServer:=grpc.NewServer()
	todo.RegisterTodoServiceServer(grpcServer, &s)
	if err:=grpcServer.Serve(lis)
		err != nil {
		log.Fatalf("Failed to serve: %v",err)
	}else {
		log.Println("Server started Successfully")
	}
		
	}
