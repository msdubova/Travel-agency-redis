package tour

import (
	"context"
	"time"
	"travel-agency-redis/internal/cache"
)

type storage interface {
	GetTours() []Tour
}

type Service struct {
	s     storage
	cache *cache.RedisCache
}

func NewService(s storage, cache *cache.RedisCache) *Service {
	service := &Service{
		s:     s,
		cache: cache,
	}

	return service
}

func (s *Service) GetTours() []Tour {

	var tours []Tour
	ctx := context.Background()
	cacheKey := "tours"

	err := s.cache.Get(ctx, cacheKey, &tours)
	if err == nil && len(tours) > 0 {
		return tours
	}

	tours = s.s.GetTours()
	s.cache.Set(ctx, cacheKey, tours, 15*time.Minute)

	return tours
}
