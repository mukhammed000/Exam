package main

import (
	"learning/service"
	"learning/storage/postgres"
	"log"
	"net"

	pb "learning/genproto/learning"

	"google.golang.org/grpc"
)

func main() {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to database", err)
	}

	log.Println("Database connected successfully! ")

	lis, err := net.Listen("tcp", ":8086")
	if err != nil {
		log.Fatal("Error while creating TCP listener: ", err.Error())
	}

	server := grpc.NewServer()

	pb.RegisterLearningServiceServer(server, service.NewLearningService(stg))

	log.Printf("Server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
