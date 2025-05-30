package kafka

import "adapter-service/constants"

type Producer interface {
	InitKafkaWriter(brokerURL, topic string)
	PublishMessage(payload string) error
}

type Consumer interface {
	StartWorker(brokerURL, topic, groupId string)
}

type kafkaQueue struct {
	queueProducer Producer
	queueConsumer Consumer
}

func NewKafkaQueue(queueProducer Producer, queueConsumer Consumer) *kafkaQueue {
	return &kafkaQueue{
		queueProducer: queueProducer,
		queueConsumer: queueConsumer,
	}
}

func (k *kafkaQueue) StartKafkaQueue(brokerURL string) {
	k.queueProducer.InitKafkaWriter(brokerURL, constants.TOPIC_JOBS_EMAIL)

	go k.queueConsumer.StartWorker(brokerURL, constants.TOPIC_JOBS_EMAIL, constants.GROUP_EMAIL_GROUP)
}
