package booking

import (
	"context"
	"math/rand"
	"strconv"
	"time"
	"travel-agency-redis/internal/cache"
)

type storage interface {
	Create(b Booking)
	GetBookings() []Booking
}

type Service struct {
	s     storage
	cache *cache.RedisCache
}

func NewService(s storage, cache *cache.RedisCache) *Service {
	return &Service{
		s:     s,
		cache: cache,
	}
}

func (s *Service) CreateBooking(b Booking) Booking {
	b.ReservationNumber = strconv.Itoa(rand.Intn(100000))
	b.Status = "pending"
	s.s.Create(b)
	s.cache.Set(context.Background(), "booking:"+b.ReservationNumber, b, 15*time.Minute)
	return b
}

func (s *Service) GetBookings() []Booking {
	var bookings []Booking
	ctx := context.Background()
	cacheKey := "bookings"

	err := s.cache.Get(ctx, cacheKey, &bookings)
	if err == nil && len(bookings) > 0 {
		return bookings
	}

	bookings = s.s.GetBookings()
	s.cache.Set(ctx, cacheKey, bookings, 15*time.Minute)

	return bookings
}
