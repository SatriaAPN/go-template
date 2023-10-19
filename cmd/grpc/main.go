package main

import (
	"fmt"
	servergrpc "go-template/server/grpc"
	"net"
)

func main() {
	// setup the network
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	srv := servergrpc.NewServer()

	go func() {
		fmt.Println("gRPC server is running on port 50051")
		if err := srv.Serve(listener); err != nil {
			fmt.Println("Failed to serve:", err)
		}
	}()

	fmt.Println("gracefully shutdown")
	srv.GracefulStop()
}
