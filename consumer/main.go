package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Define Kafka reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("KAFKA_BROKER_URL")},
		GroupID:   "consumer-group-1",
		Topic:     "myTopic",
	})
	defer reader.Close()

	conn, err := net.Dial("tcp", os.Getenv("SOCKET_HOST"))
	if err != nil {
		log.Fatal("Error connecting to socket:", err)
	}
	defer conn.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error reading message:", err)
		}
		log.Printf("Received: %s\n", string(msg.Value))

		_, err = conn.Write(msg.Value) // Write to the socket channel
		if err != nil {
			log.Fatal("Error sending to socket:", err)
		}
	}
}
