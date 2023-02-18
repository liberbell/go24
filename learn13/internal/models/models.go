package models

import "time"

type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
