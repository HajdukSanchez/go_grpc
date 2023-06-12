package main

import (
	"log"
	"net"

	"github.com/hajduksanchez/go_grpc/database"
	"github.com/hajduksanchez/go_grpc/proto/testpb"
	"github.com/hajduksanchez/go_grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5070")
	if err != nil {
		log.Fatal(err)
	}

	// Create new repository instance
	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Create new server instance for student
	server := server.NewTestServer(repo)

	// GRPC server
	grpcServer := grpc.NewServer()
	testpb.RegisterTestServiceServer(grpcServer, server)

	// Reflection to provide some meta-data for clients
	reflection.Register(grpcServer)

	// Validate if there is no error setting up the app and his services
	if err := grpcServer.Serve(list); err != nil {
		log.Fatal(err)
	}
}
