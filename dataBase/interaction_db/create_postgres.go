package interaction_db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// создание таблицы
func CreateBookTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	author VARCHAR(100) NOT NULL,
	review VARCHAR(1000),
	published_year INTEGER NOT NULL,
	read BOOLEAN NOT NULL,
	receiving_at TIMESTAMP NOT NULL,
	read_at TIMESTAMP,

	UNIQUE (title)
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
