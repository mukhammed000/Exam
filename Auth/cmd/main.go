package main

import (
    "log"
    "net"
    "users/service"
    "users/storage/postgres"
    "users/kafka" // Import your Kafka package

    pb "users/genproto/users"

    "google.golang.org/grpc"
)

func main() {
    stg, err := postgres.NewPostgresStorage()
    if err != nil {
        log.Fatal("Error while connecting to database", err)
    }

    log.Println("Database connected successfully!")

    kafkaProducer, err := kafka.NewKafkaProducer([]string{"localhost:9092"})
    if err != nil {
        log.Fatal("Error while creating Kafka producer", err)
    }
    defer kafkaProducer.Close()

    lis, err := net.Listen("tcp", ":8084")
    if err != nil {
        log.Fatal("Error while creating TCP listener: ", err.Error())
    }

    server := grpc.NewServer()

    userService := service.NewUserService(stg, kafkaProducer)
    pb.RegisterUsersServiceServer(server, userService)

    log.Printf("Server listening at %v", lis.Addr())

    if err := server.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
