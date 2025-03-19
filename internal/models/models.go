package models

import (
	"gorm.io/gorm"
	"time"
)

type StockAPIResponse struct {
	Items    []StockItemSerialize `json:"items"`
	NextPage string               `json:"next_page"`
}

type StockItemSerialize struct {
	Ticker     string    `json:"ticker"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Company    string    `json:"company"`
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	Time       time.Time `json:"time"`
}

type StockItem struct {
	gorm.Model
	Ticker     string
	TargetFrom float64
	TargetTo   float64
	Company    string
	Action     string
	Brokerage  string
	RatingFrom string
	RatingTo   string
	Time       time.Time `gorm:"index"`
}
