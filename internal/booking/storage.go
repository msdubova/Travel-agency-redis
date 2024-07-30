package booking

import (
	"sync"
)

type InMemStorage struct {
	bookingsM sync.Mutex
	bookings  []Booking
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{}
}

func (s *InMemStorage) Create(b Booking) {
	s.bookingsM.Lock()
	defer s.bookingsM.Unlock()

	s.bookings = append(s.bookings, b)
}

func (s *InMemStorage) GetBookings() []Booking {
	s.bookingsM.Lock()
	defer s.bookingsM.Unlock()

	return s.bookings
}
