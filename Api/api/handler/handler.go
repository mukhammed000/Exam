package handler

import (
	pb "api/genproto"
	"api/kafka"

	"github.com/go-redis/redis/v8"
)

type Handler struct {
	Auth          pb.UsersServiceClient
	Learning      pb.LearningServiceClient
	Progress      pb.ProgressServiceClient
	rdb           *redis.Client
	kafkaProducer *kafka.KafkaProducer
}

func NewHandler(auth pb.UsersServiceClient, pro pb.ProgressServiceClient, learn pb.LearningServiceClient, rdb *redis.Client, kafkaProducer *kafka.KafkaProducer) *Handler {
	return &Handler{
		Auth:          auth,
		Progress:      pro,
		Learning:      learn,
		rdb:           rdb,
		kafkaProducer: kafkaProducer,
	}
}
