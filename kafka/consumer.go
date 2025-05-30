package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type consumer struct{}

func NewQueueConsumerImpl() Consumer {
	return &consumer{}
}

// ========================================
// # E-MAIL
// ========================================
func (c *consumer) StartWorker(brokerURL, topic, groupId string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerURL},
		GroupID: groupId,
		Topic:   topic,
	})
	logrus.Infoln("Email worker started...")
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logrus.Fatalf("ReadMessage error: %s", err.Error())
			continue
		}
		logrus.Infof("📬 Received job: %s\n", string(m.Value))
	}
}
