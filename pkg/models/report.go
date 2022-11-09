package models

type GetTransactions struct {
	User         int64  `json:"user"`
	DateFrom     string `json:"dateFrom"`
	DateTo       string `json:"dateTo"`
	SortBy       string `json:"sortBy"`
	SortOrder    string `json:"sortOrder"`
	NumberOfPage int    `json:"page"`
}

type GetService struct {
	Year  string `json:"year"`
	Month string `json:"month"`
}
