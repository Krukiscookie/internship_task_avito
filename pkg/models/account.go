package models

type Account struct {
	Id      int     `json:"id"`
	Balance float64 `json:"balance"`
	Reserve float64 `json:"reserve"`
}

type SwagAccount struct {
	Id      int     `json:"id"`
	Balance float64 `json:"balance"`
}
