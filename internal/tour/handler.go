package tour

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

var s InMemStorage

type service interface {
	GetTours() []Tour
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{
		s: s,
	}
}

func (h Handler) GetTours(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tours := h.s.GetTours()

	err := json.NewEncoder(w).Encode(tours)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Debug().Err(err).Msg("Failed to encode JSON response")
		return
	}

}
