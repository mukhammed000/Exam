package main

import (
	"fmt"
	"log"

	"api/api"
	"api/api/handler"
	pb "api/genproto"
	"api/kafka" // Import your Kafka package

	_ "api/docs"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up gRPC connections
	userConn, err := grpc.Dial("localhost:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while creating new client: ", err.Error())
	}
	defer userConn.Close()

	progressConn, err := grpc.Dial("localhost:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while creating new client: ", err.Error())
	}
	defer progressConn.Close()

	learningConn, err := grpc.Dial("localhost:8086", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while creating new client: ", err.Error())
	}
	defer learningConn.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	kafkaProducer, err := kafka.NewKafkaProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatal("Error while creating Kafka producer: ", err.Error())
	}
	defer kafkaProducer.Close()

	usc := pb.NewUsersServiceClient(userConn)
	pro := pb.NewProgressServiceClient(progressConn)
	lsc := pb.NewLearningServiceClient(learningConn)

	h := handler.NewHandler(usc, pro, lsc, rdb, kafkaProducer)

	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
