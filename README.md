### Kafka Payment Processor Example
This example demonstrates a real-time payment processor example using **Kafka** to ingest transactions, process them, and detect potential fraud. Kafka topics are used to communicate between microservices, where a consumer service listens for transactions and flags fraudulent activity.

---

### Prerequisites

Before starting the Kafka fraud detection example, ensure you have the following software installed:

- [Go](https://golang.org/doc/install) (1.16 or above)
- [Docker](https://docs.docker.com/get-docker/)

---

### 1. Set up Kafka and Zookeeper using Docker

Create a `docker-compose.yml` file to configure **Kafka** and **Zookeeper**.

#### Docker Compose Setup

1. Create the following `docker-compose.yml` file in the root directory of the project:

    ```yaml
    version: '3'
    services:
      zookeeper:
        image: confluentinc/cp-zookeeper:7.3.2
        container_name: zookeeper
        ports:
          - "2181:2181"
        environment:
          ZOOKEEPER_CLIENT_PORT: 2181
          ZOOKEEPER_TICK_TIME: 2000

      kafka:
        image: confluentinc/cp-kafka:7.3.2
        container_name: kafka
        ports:
          - "9092:9092"
        environment:
          KAFKA_BROKER_ID: 1
          KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
          KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
          KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT
          KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
          KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
        depends_on:
          - zookeeper
    ```

2. To start Kafka and Zookeeper, run:

    ```bash
    docker-compose up -d
    ```

   This will launch Kafka on port `9092` and Zookeeper on port `2181`.

3. **Verify** Kafka and Zookeeper are running:

    ```bash
    docker ps
    ```

4. Create Kafka topics to handle fraud detection, such as `transactions-topic`:

    ```bash
    docker exec -it kafka kafka-topics --create --topic transactions-topic --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1
    ```

---

### 2. Application Setup

#### Directory Structure

The structure of the Kafka Fraud Detection application might look like this:

```bash
kafka-fraud-detection/
│
├── cmd/
│   ├── producer/
│   ├   └── main.go
│   └── consumer/
│   ├   └── main.go
├── internal/
│   ├── config/
│   ├   └── config.go
│   ├── kafka/
│   ├   └── consumer.go
│   ├   └── producer.go
│   └── fraud/
│   ├   └── detection.go
│   ├── payment/
│   ├   └── event.go
├── docker-compose.yml
├── go.mod
└── README.md
```

#### Producer Service

The producer simulates transaction data and pushes it to the Kafka topic. The producer is placed under `cmd/producer`.

#### Consumer Service

The consumer reads transactions from the Kafka topic, processes them, and identifies potential fraud. The fraud detection logic is handled in `pkg/fraud` and consumed in `cmd/consumer`.

---

### 3. Running the Producer and Consumer

#### Env config
Create a .env and set the below

```
KAFKA_BROKER=localhost:9092
KAFKA_TOPIC=payments-topic
```


#### Producer

To start the producer, which sends transactions to Kafka:

```bash
go run cmd/producer/main.go
```

#### Consumer

To start the consumer, which listens for transactions and detects fraud:

```bash
go run cmd/consumer/main.go
```

---

### 4. Stopping the Services

To stop Kafka and Zookeeper, run:

```bash
docker-compose down
```

---

### Example Usage

- **Producer**: Sends simulated transactions to the `payments-topic` configured in the .env
- **Consumer**: Consumes these transactions and checks for fraud by comparing each transaction against predefined rules.

### Notes

- The consumer service can be extended to include fraud detection logic like IP analysis, geolocation mismatches, or unusual spending patterns.
- Both the producer and consumer leverage **Kafka** for real-time processing of transactions, making this architecture suitable for large-scale financial services.

---
