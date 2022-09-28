package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/Akashkumar-Jeyaramans/snmpOperations/protos/v1/commands"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

const (
	port = ":9090"
)

type snmpgrpcServer struct {
	pb.UnimplementedCommandServer
}

func (s *snmpgrpcServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Message: fmt.Sprintf("name: %s", in.IpAddress),
	}, nil
}

func main() {

	// go func() {
	// 	//mux
	// 	mux := runtime.NewServeMux()
	// 	//registerhandle
	// 	pb.RegisterCommandHandlerServer(context.Background(), mux, &snmpgrpcServer{})
	// 	//httpserver
	// 	lis := http.ListenAndServe("localhost:8081", mux)
	// 	log.Fatalln("succes", lis)
	// }()

	//runGatewayServer()

	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux()
	// setting up a dail up for gRPC service by specifying endpoint/target url
	//err := pb.RegisterCommandHandlerFromEndpoint(context.Background(), mux, "localhost:8080", []grpc.DialOption{grpc.WithInsecure()})
	err := pb.RegisterCommandHandlerServer(context.Background(), mux, &snmpgrpcServer{})
	if err != nil {
		log.Fatal(err)
	}
	// Creating a normal HTTP server
	server := http.Server{
		Handler: mux,
	}
	// creating a listener for server
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	// start server
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}

}

func runGatewayServer() {

	// 	grpcMux := runtime.NewServeMux()

	// 	ctx, cancel := context.WithCancel(context.Background())
	// 	defer cancel()

	// 	err := pb.RegisterCommandHandlerServer(ctx, grpcMux, &snmpgrpcServer{})
	// 	if err != nil {
	// 		log.Fatal("cannot register handler server:", err)
	// 	}

	// 	mux := http.NewServeMux()
	// 	mux.Handle("/", grpcMux)

	// 	lis, err := http.Listen("tcp", "9090")
	// 	if err != nil {
	// 		log.Fatal("cannot create listener:", err)
	// 	}

	// 	log.Printf("start HTTP gateway server at %s", lis.Addr().String())
	// 	err = http.Serve(lis, mux)
	// 	if err != nil {
	// 		log.Fatal("cannot start HTTP gateway server:", err)
	// 	}

	//mux
	// mux := runtime.NewServeMux()
	// //registerhandle
	// pb.RegisterCommandHandlerServer(context.Background(), mux, &snmpgrpcServer{})
	// //httpserver
	// log.Fatalln(http.ListenAndServe("localhost:8081", mux))

	// // creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	// mux := runtime.NewServeMux()
	// // setting up a dail up for gRPC service by specifying endpoint/target url
	// //err := pb.RegisterCommandHandlerFromEndpoint(context.Background(), mux, "localhost:8080", []grpc.DialOption{grpc.WithInsecure()})
	// err := pb.RegisterCommandHandlerServer(context.Background(), mux, &snmpgrpcServer{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // Creating a normal HTTP server
	// server := http.Server{
	// 	Handler: mux,
	// }
	// // creating a listener for server
	// l, err := net.Listen("tcp", ":8081")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // start server
	// err = server.Serve(l)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
