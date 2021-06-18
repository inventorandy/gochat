package main

import (
	"fmt"
	"gochat/platform/internal/proto/pb"
	"gochat/platform/services/account/server"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	// Get the Port for the Account Service
	port, exists := os.LookupEnv("ACCOUNT_SERVICE_PORT")
	if !exists {
		log.Fatalln("No ACCOUNT_SERVICE_PORT environment variable specified. Shutting down...")
	}

	// Create a TCP Listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalln("Could not start TCP server: ", err.Error())
	}

	// Create the new Service Definition
	accountSrv, err := server.NewAccountService()
	if err != nil {
		log.Fatalln("Could not create server instance: ", err.Error())
	}

	// Create a gRPC Server and Register the Definitions
	grpcSrv := grpc.NewServer()
	pb.RegisterAccountServiceServer(grpcSrv, accountSrv)

	// Log a quick message
	log.Println("Starting Account service...")

	// Run the Server
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalln("Critical Error in Account Service: ", err.Error())
	}
}
