package main

import (
	"log"
	"sd/internal/kafka"
	"time"
)

func main() {
	for {
		err := kafka.ProduceMessage("localhost:9092", "test-topic", 1000)
		if err != nil {
			log.Fatalf("Failed to produce messages: %v", err)
			continue
		}
		time.Sleep(2 * time.Second)
	}
	log.Printf("Successfully produced 1000 messages batches to topic 'test-topic'")
}
