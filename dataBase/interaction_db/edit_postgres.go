package interaction_db

import (
	"context"
	"postgres/dataBase/interaction_db/model"

	"github.com/jackc/pgx/v5"
)

func Edit_Book(ctx context.Context, conn *pgx.Conn, book model.ModelBook) error {
	sqlQuery := `
	UPDATE books
	SET review=$1, read=$2, read_at=$3
	WHERE id=$4 
	`

	_, err := conn.Exec(ctx, sqlQuery, book.Review, book.Read, book.Read_at, book.ID)

	return err
}
