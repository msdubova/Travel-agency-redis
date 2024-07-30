package tour

import (
	"sync"
	"time"
)

type InMemStorage struct {
	toursM sync.Mutex
	tours  []Tour
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		tours: []Tour{
			{ID: "1", Destination: "Spain, Alicante", Description: "Enjoy a sunny beach", Price: 299.99, StartDate: time.Now().Add(48 * time.Hour), EndDate: time.Now().Add(7 * 24 * time.Hour)},
			{ID: "2", Destination: "Austria, Courshavel", Description: "Explore the mountains", Price: 399.99, StartDate: time.Now().Add(96 * time.Hour), EndDate: time.Now().Add(10 * 24 * time.Hour)},
		},
	}
}

func (s *InMemStorage) GetTours() []Tour {
	s.toursM.Lock()
	defer s.toursM.Unlock()

	toursCopy := make([]Tour, len(s.tours))
	copy(toursCopy, s.tours)

	return toursCopy
}
