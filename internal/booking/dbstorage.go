package booking

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) *PostgresStorage {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) Create(b Booking) {
	_, err := s.db.Exec(
		"INSERT INTO bookings (reservation_number, name, email, price, status, tour_id) VALUES ($1, $2, $3, $4, $5, $6)",
		b.ReservationNumber, b.Name, b.Email, b.Price, b.Status, b.TourID,
	)
	if err != nil {
		log.Println("Failed to create booking:", err)
	}
}

func (s *PostgresStorage) GetBookings() []Booking {
	rows, err := s.db.Query("SELECT reservation_number, name, email, price, status, tour_id FROM bookings")
	if err != nil {
		log.Println("Failed to fetch bookings:", err)
		return nil
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var b Booking
		if err := rows.Scan(&b.ReservationNumber, &b.Name, &b.Email, &b.Price, &b.Status, &b.TourID); err != nil {
			log.Println("Failed to scan booking:", err)
			continue
		}
		bookings = append(bookings, b)
	}

	return bookings
}
