package model

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	State   string `json:"state"`
}
