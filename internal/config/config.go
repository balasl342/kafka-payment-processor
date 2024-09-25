package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Kafka KafkaConfig
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

// LoadConfig loads configurations from environment variables or default values
func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	return Config{
		Kafka: KafkaConfig{
			Brokers: []string{os.Getenv("KAFKA_BROKER")},
			Topic:   os.Getenv("KAFKA_TOPIC"),
		},
	}
}
