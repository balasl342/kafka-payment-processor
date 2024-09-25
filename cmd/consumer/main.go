package main

import (
	"log"

	"github.com/balasl342/kafka-payment-processor/internal/config"
	"github.com/balasl342/kafka-payment-processor/internal/kafka"
)

func main() {
	cfg := config.LoadConfig() // Load Kafka brokers, topics, etc.

	consumer, err := kafka.NewConsumer(cfg.Kafka.Topic)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	defer consumer.Close()

	kafka.ProcessPaymentEvents(consumer, cfg.Kafka.Topic) // Start processing events
}
