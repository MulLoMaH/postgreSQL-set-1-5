package interaction_db

import (
	"context"
	"fmt"
	"postgres/dataBase/interaction_db/model"

	"github.com/jackc/pgx/v5"
)

func List_Pages(ctx context.Context, conn *pgx.Conn, N int) error {
	nStr := (N - 1) * 3

	sqlQuery := `
	SELECT *
	FROM books
	LIMIT 3
	OFFSET $1
	`

	rows, err := conn.Query(ctx, sqlQuery, nStr)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {

		var book model.ModelBook
		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Review,
			&book.Published_year,
			&book.Read,
			&book.Receiving_at,
			&book.Read_at,
		); err != nil {
			return err
		}

		print_Book(book)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func print_Book(book model.ModelBook) {
	fmt.Println("-----------------------------")
	fmt.Println("ID :", book.ID)
	fmt.Println("Title: ", book.Title)
	fmt.Println("Author: ", book.Author)
	if book.Review.Valid {
		fmt.Println("Review: ", book.Review.String)
	} else {
		fmt.Println("Review: (нет отзыва)")
	}
	fmt.Println("Published year", book.Published_year)
	fmt.Println("Read", book.Read)
	fmt.Println("Receiving at", book.Receiving_at)
	if book.Read_at != nil {
		fmt.Println("Read at", book.Read_at)
	} else {
		fmt.Println("Read at: (книга не прочитана)")
	}
}
