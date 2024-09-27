package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/balasl342/kafka-payment-processor/internal/payment"
)

func NewKafkaProducer(broker string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
	return producer
}

func ProducePaymentEvent(producer sarama.SyncProducer, topic string, event payment.PaymentEvent) {
	messageBytes, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("Error marshaling payment event: %v", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageBytes),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Error sending message to Kafka: %v", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
