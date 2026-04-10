package interaction_db

import (
	"context"
	"fmt"
	"postgres/dataBase/interaction_db/model"

	"github.com/jackc/pgx/v5"
)

func Get_All_Books(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	SELECT id, title, author, review, published_year, read, receiving_at, read_at
	FROM books;
	`

	var books []model.ModelBook

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {

		if err = rows.Err(); err != nil {
			return err
		}

		var book model.ModelBook

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Review,
			&book.Published_year,
			&book.Read,
			&book.Receiving_at,
			&book.Read_at,
		)
		if err != nil {
			return err
		}

		books = append(books, book)

		print_Books(book)
	}

	return nil
}

func print_Books(book model.ModelBook) {
	fmt.Println("-----------------------------")
	fmt.Println("ID :", book.ID)
	fmt.Println("Title: ", book.Title)
	fmt.Println("Author: ", book.Author)
	if book.Review.Valid {
		fmt.Println("Review: ", book.Review)
	} else {
		fmt.Println("Review: (нет отзыва)")
	}
	fmt.Println("Published year", book.Published_year)
	fmt.Println("Read", book.Read)
	fmt.Println("Receiving at", book.Receiving_at)
	fmt.Println("Read at", book.Read_at)
}
