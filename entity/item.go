package entity

import (
	"time"
)

type Item struct {
	ItemId      int
	ItemCode    string
	Quantity    int
	Description string
	OrderId     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
