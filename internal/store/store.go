package store

import (
	"sd/internal/model"
	"sync"
)

type Store struct{
	mu sync.Mutex
	Data map[string][]model.Message
	Total int
}

func NewStore() *Store {
	return &Store{
		Data:map[string][]model.Message{
          "completed": {},
		  "failed":    {},
		  "pending":   {},
		},
	}
}

func( s *Store) AddMessage(msg model.Message){
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Data[msg.State] = append(s.Data[msg.State], msg)
	s.Total++
}

func (s *Store) GetMessagesCount() int{
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Total
}

func (s *Store) GetBucketCount(state string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	if messages, exists := s.Data[state]; exists {
		return len(messages)
	}
	return 0
}


func (s *Store) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Data = map[string][]model.Message{
		"completed": {},
		"failed":    {},
		"pending":   {},
	}
	s.Total = 0
}
