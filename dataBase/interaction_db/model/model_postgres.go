package model

import (
	"database/sql"
	"time"
)

type ModelBook struct {
	ID             int
	Title          string
	Author         string
	Review         sql.NullString
	Published_year int
	Read           bool
	Receiving_at   time.Time
	Read_at        *time.Time
}
