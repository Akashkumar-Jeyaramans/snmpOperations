package main

import (
	"log"
	"net"

	pb "github.com/Akashkumar-Jeyaramans/snmpOperations/v1"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type snmpgrpcServer struct {
	pb.UnimplementedCommandServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommandServer(s, &snmpgrpcServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen %v", err)
	}
}
