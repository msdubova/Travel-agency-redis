package booking

type Booking struct {
	ReservationNumber string  `json:"reservnumber"`
	Name              string  `json:"name"`
	Email             string  `json:"email"`
	Price             float64 `json:"price"`
	Status            string  `json:"status"`
	TourID            string  `json:"tourId"`
}

type Bookings []Booking
