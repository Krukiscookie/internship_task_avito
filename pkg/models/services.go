package models

import "time"

type Services struct {
	Id        int       `json:"id"`
	AccountId int       `json:"account"`
	IdService int       `json:"id-service"`
	IdOrder   int       `json:"id-order"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	TimeStamp time.Time `json:"timestamp"`
}

type SwagServices struct {
	AccountId int     `json:"account"`
	IdService int     `json:"id-service"`
	IdOrder   int     `json:"id-order"`
	Amount    float64 `json:"amount"`
}
