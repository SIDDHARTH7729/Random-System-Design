package main

import (
	"log"
	"sd/internal/kafka"
	"sd/internal/store"
)


func main() {
	s := store.NewStore()
	err := kafka.StartConsumer("localhost:9092", "test-topic", s)
	if err != nil{
        log.Fatalf("Failed to start consumer: %v", err)
		return
	}
}