package models

import (
	"time"
)

type Author struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Biography string    `json:"biography"`
	BirthDate time.Time `json:"birth_date"`
}
