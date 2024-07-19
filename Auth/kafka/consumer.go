package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
)

type KafkaConsumer struct {
	consumerGroup sarama.ConsumerGroup
	redisClient   *redis.Client
}

func NewKafkaConsumer(brokers []string, groupId string, redisClient *redis.Client) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Session.Timeout = 10 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	// Additional configuration options can be set here

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupId, config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumerGroup: consumerGroup,
		redisClient:   redisClient,
	}, nil
}

func (kc *KafkaConsumer) Start(ctx context.Context) {
	go func() {
		for {
			if err := kc.consumerGroup.Consume(ctx, []string{"reset_password_topic"}, kc); err != nil {
				log.Fatal("Error while consuming messages: ", err)
			}
		}
	}()
}

func (kc *KafkaConsumer) Close() error {
	if err := kc.consumerGroup.Close(); err != nil {
		return err
	}
	return kc.redisClient.Close()
}

func (kc *KafkaConsumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (kc *KafkaConsumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (kc *KafkaConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Received message: %s", string(message.Value))

		if err := kc.processMessage(string(message.Value)); err != nil {
			log.Printf("Error processing message: %s", err)
		}

		sess.MarkMessage(message, "")
	}
	return nil
}

func (kc *KafkaConsumer) processMessage(message string) error {
	var email, code string
	if _, err := fmt.Sscanf(message, "Reset code for %s is %s", &email, &code); err != nil {
		return err
	}

	expiration := 15 * time.Minute
	key := fmt.Sprintf("password_code:%s", email)
	if err := kc.redisClient.Set(context.Background(), key, code, expiration).Err(); err != nil {
		return err
	}

	return nil
}
