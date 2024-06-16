package models

import (
	"time"
)

type Delivery struct {
	CustomerName	string    `json:"customer_name"`
	Address			string    `json:"address"`
	Distance		float64   `json:"distance"`
	ReceivedAt		time.Time
	ID				int		
}

type Courier struct {
	ID	int
}