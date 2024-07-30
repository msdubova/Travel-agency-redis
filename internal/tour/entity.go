package tour

import "time"

type Tour struct {
	ID          string    `json:"id"`
	Price       float64   `json:"price"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Destination string    `json:"destination"`
	Description string    `json:"description"`
	PrimePrice  float64
}

type Tours []Tour
