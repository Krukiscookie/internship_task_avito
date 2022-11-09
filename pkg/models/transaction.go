package models

import "time"

type Transaction struct {
	Id        int       `json:"id"`
	IdFrom    int       `json:"from"`
	IdTo      int       `json:"to"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	TimeStamp time.Time `json:"timestamp"`
}

type SwagTransaction struct {
	IdFrom int     `json:"from"`
	IdTo   int     `json:"to"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}
