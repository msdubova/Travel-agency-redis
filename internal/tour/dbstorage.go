package tour

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

func (s *PostgresStorage) GetTours() []Tour {
	rows, err := s.db.Query("SELECT id, price, start_date, end_date, destination, description FROM tours")
	if err != nil {
		log.Println("Failed to fetch tours:", err)
		return nil
	}
	defer rows.Close()

	var tours []Tour
	for rows.Next() {
		var t Tour
		if err := rows.Scan(&t.ID, &t.Price, &t.StartDate, &t.EndDate, &t.Destination, &t.Description); err != nil {
			log.Println("Failed to scan tour:", err)
			continue
		}
		tours = append(tours, t)
	}

	return tours
}
