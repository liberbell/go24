package dbrepo

import "github.com/tsawler/bookings/internal/models"

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {
	stmt := `INSERT INTO reservation (first_name, last_name, emain, phone, start_date, end_data, room_id, created_at, updated_at)
	         values ($1, $2, $3, $4, $5, $6, $7)`
	return nil
}
