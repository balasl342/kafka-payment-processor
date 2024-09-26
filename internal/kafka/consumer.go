package kafka

import (
	"encoding/json"
	"log"

	"github.com/balasl342/kafka-payment-processor/internal/fraud"
	"github.com/balasl342/kafka-payment-processor/internal/payment"

	"github.com/IBM/sarama"
)

func NewConsumer(topic string) (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func ProcessPaymentEvents(consumer sarama.Consumer, topic string) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start consumer for topic '%s': %v", topic, err)
	}

	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		var event payment.PaymentEvent
		err := json.Unmarshal(message.Value, &event)
		if err != nil {
			log.Printf("Failed to unmarshal event: %v", err)
			continue
		}

		log.Printf("Processing payment event: %+v", event)

		if fraud.IsFraudulent(event) {
			log.Printf("Fraudulent transaction detected! Transaction ID: %s", event.TransactionID)
			continue
		}

		log.Printf("Transaction %s processed successfully", event.TransactionID)
	}
}
