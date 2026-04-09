package interaction_db

import (
	"context"
	"postgres/dataBase/interaction_db/model"

	"github.com/jackc/pgx/v5"
)

// создание нового экземпляра книги
func Insert_book(ctx context.Context, conn *pgx.Conn, book model.ModelBook) error {
	sqlQuery := `
	INSERT INTO books (
	title,
	author,
	published_year,
	read,
	receiving_at
)
	VALUES ($1, $2, $3, $4, $5);
	`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		book.Title,
		book.Author,
		book.Published_year,
		book.Read,
		book.Receiving_at,
	)

	return err
}
