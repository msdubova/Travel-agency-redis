package tour

import "time"

type storage interface {
	GetTours() []Tour
}

type Service struct {
	s storage
}

func NewService(s storage) *Service {
	service := &Service{
		s: s,
	}

	return service
}

func (s *Service) GetTours() []Tour {

	tours := s.s.GetTours()
	for i, tour := range tours {
		tours[i].PrimePrice = tour.Price * 1.1
	}
	now := time.Now()
	isEvening := now.Hour() >= 18 || now.Hour() < 22

	if isEvening {
		for i, tour := range tours {
			tours[i].Price = tour.PrimePrice
		}
	}

	return tours
}
