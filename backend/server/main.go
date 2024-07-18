package main

import (
	"log"
	"net"

	"github.com/ultravioletasdf/messenger/backend/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln("Failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterUsersServer(s, &usersServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to server:", err)
	}
}
