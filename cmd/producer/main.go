package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/balasl342/kafka-payment-processor/internal/kafka"
	"github.com/balasl342/kafka-payment-processor/internal/payment"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	producer := kafka.NewKafkaProducer(broker)
	defer producer.Close()

	// Payment event production
	for i := 1; i <= 10; i++ {
		event := payment.PaymentEvent{
			TransactionID: "txn" + strconv.Itoa(i),
			Amount:        float64(5000 * i),
		}
		kafka.ProducePaymentEvent(producer, topic, event)
		time.Sleep(2 * time.Second)
	}
}
