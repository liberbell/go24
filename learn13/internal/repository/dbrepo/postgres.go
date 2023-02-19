package dbrepo

import "github.com/tsawler/bookings/internal/models"

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {
	return nil
}
