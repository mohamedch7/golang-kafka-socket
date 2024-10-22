package main

import (
    "context"
    "fmt"
    "github.com/segmentio/kafka-go"
    "log"
    "os"
)

func main() {
    // Define Kafka writer
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  []string{os.Getenv("KAFKA_BROKER_URL")},
        Topic:    "myTopic",
        Balancer: &kafka.LeastBytes{},
    })
    defer writer.Close()

    messages := []kafka.Message{
        {Value: []byte("message1")},
        {Value: []byte("message2")},
    }

    err := writer.WriteMessages(context.Background(), messages...)
    if err != nil {
        log.Fatalf("could not write message: %v", err)
    }

    fmt.Println("Produced messages to Kafka!")
}
