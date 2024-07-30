package booking

import (
	"math/rand"
	"strconv"
)

type storage interface {
	Create(b Booking)
	GetBookings() []Booking
}

type Service struct {
	s storage
}

func NewService(s storage) *Service {
	return &Service{
		s: s,
	}
}

func (s *Service) CreateBooking(b Booking) Booking {
	b.ReservationNumber = strconv.Itoa(rand.Intn(100000))
	b.Status = "pending"
	s.s.Create(b)
	return b
}

func (s *Service) GetBookings() []Booking {
	return s.s.GetBookings()
}
