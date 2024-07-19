package kafka

import (
	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{producer: producer}, nil
}

// SendMessage sends a Kafka message.
func (p *KafkaProducer) SendMessage(msg *sarama.ProducerMessage) error {
	_, _, err := p.producer.SendMessage(msg)
	return err
}

func (p *KafkaProducer) Close() error {
	return p.producer.Close()
}
