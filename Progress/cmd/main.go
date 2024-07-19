package main

import (
	"log"
	"net"
	"progress/storage/postgres"

	"progress/service"

	pb "progress/genproto/progress"

	"google.golang.org/grpc"
)

func main() {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to database", err)
	}

	log.Println("Database connected successfully! ")

	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal("Error while creating TCP listener: ", err.Error())
	}

	server := grpc.NewServer()

	pb.RegisterProgressServiceServer(server, service.NewProgressService(stg))
	
	log.Printf("Server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
