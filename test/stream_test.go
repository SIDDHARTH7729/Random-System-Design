package test

import (
	"context"
	"log"
	"sd/internal/model"
	"sd/internal/store"
	"testing"
	"time"
)

func TestNoMessagesDropped(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := store.NewStore()
	totalMessages := 1000
	loop := 0

	for {
		select {
		case <-ctx.Done():
			log.Println("Received interrupt signal, stopping test...")
			return
		default:
			// Clear the store for a fresh test in each loop (simulate independent cycles)
			s.Reset()

			seenIds := make(map[string]bool)

			for i := 0; i < totalMessages; i++ {
				// Use unique ID per loop
				id := "msg-" + string(rune(loop)) + "-" + string(rune(i))
				state := "completed"
				if i%3 == 1 {
					state = "failed"
				} else if i%3 == 2 {
					state = "pending"
				}

				if seenIds[id] {
					t.Errorf("Duplicate message ID found: %s", id)
				}
				seenIds[id] = true

				s.AddMessage(model.Message{
					ID:      id,
					Content: "This is message " + id,
					State:   state,
				})
			}

			count := s.GetMessagesCount()
			if count != totalMessages {
				t.Errorf("Expected %d messages, but got %d", totalMessages, count)
			}

			states := []string{"completed", "failed", "pending"}
			sum := 0
			for _, state := range states {
				count := s.GetBucketCount(state)
				t.Logf("State: %s, Count: %d", state, count)
				sum += count
			}

			if sum != totalMessages {
				t.Errorf("Total messages in buckets %d, expected %d", sum, totalMessages)
			}

			loop++
			time.Sleep(2 * time.Second)
		}
	}
}
