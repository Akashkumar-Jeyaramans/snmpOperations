package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/Akashkumar-Jeyaramans/snmpOperations/v1/commands"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type getServer struct {
	pb.UnimplementedCommandServer
}

func (s getServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Message: in.IpAddress,
	}, nil
}

func main() {
	mux := runtime.NewServeMux()

	err := pb.RegisterCommandHandlerServer(context.Background(), mux, &getServer{})
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: mux,
	}
	// creating a listener for server
	l, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("server listening : %v", l)
	}
	// start server
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}

}
