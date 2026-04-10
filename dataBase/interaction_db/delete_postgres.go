package interaction_db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Delete_Books(ctx context.Context, conn *pgx.Conn, booksIDs []int) error {
	sqlQuery := `
		DELETE FROM books 
		WHERE id = ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, booksIDs)

	return err
}
