package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type producer struct{}

func NewQueueProducerImpl() Producer {
	return &producer{}
}

var Writer *kafka.Writer

func (p *producer) InitKafkaWriter(brokerURL, topic string) {
	Writer = &kafka.Writer{
		Addr:     kafka.TCP(brokerURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		// BatchSize:    10,
		// BatchTimeout: 1 * time.Second,
		Async: true,
	}
}

func (p *producer) PublishMessage(payload string) error {
	msg := kafka.Message{
		Key:   []byte(time.Now().Format(time.RFC3339)),
		Value: []byte(payload),
	}
	return Writer.WriteMessages(context.Background(), msg)

}
