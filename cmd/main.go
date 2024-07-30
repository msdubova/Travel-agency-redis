package main

import (
	"Travel-agency/internal/booking"
	"Travel-agency/internal/tour"
	"os"

	"github.com/rs/zerolog/log"

	"net/http"
)

func main() {
	mux := http.NewServeMux()
	connStr := os.Getenv("POSTGRES_CONN_STR")

	// tourStorage := tour.NewInMemStorage()
	// tourService := tour.NewService(tourStorage)
	// tourHandler := tour.NewHandler(tourService)
	// mux.HandleFunc("GET /tours", tourHandler.GetTours)

	// bookingStorage := booking.NewInMemStorage()
	// bookingService := booking.NewService(bookingStorage)
	// bookingHandler := booking.NewHandler(bookingService)

	// mux.HandleFunc("POST /booking", bookingHandler.CreateBooking)
	// mux.HandleFunc("GET /bookings", bookingHandler.GetBookings)

	tourStorage := tour.NewPostgresStorage(connStr)
	tourService := tour.NewService(tourStorage)
	tourHandler := tour.NewHandler(tourService)
	mux.HandleFunc("GET /tours", tourHandler.GetTours)

	bookingStorage := booking.NewPostgresStorage(connStr)
	bookingService := booking.NewService(bookingStorage)
	bookingHandler := booking.NewHandler(bookingService)
	mux.HandleFunc("POST /booking", bookingHandler.CreateBooking)
	mux.HandleFunc("GET /bookings", bookingHandler.GetBookings)

	error := http.ListenAndServe(":8080", mux)
	if error != nil {
		log.Fatal().Err(error).Msg("Failed to listen and serve")
	}
}
