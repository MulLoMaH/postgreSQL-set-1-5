package main

import (
	"context"
	"log"
	"postgres/dataBase/connect_db"
	"postgres/dataBase/interaction_db"
)

func main() {
	ctx := context.Background()
	conn, err := connect_db.Connect_db(ctx) //вызов функции подключения к базе данных
	if err != nil {
		log.Fatal(err)
	}

	//вызов функции создания таблицы
	if err := interaction_db.CreateBookTable(ctx, conn); err != nil {
		log.Fatal(err)
	}

	/*
		title := []string{
			"Стол",
			"Стул",
			"Пирог",
			"Волна",
			"Арбуз",
			"Дыня",
			"Помидор",
			"Игорь",
			"Виноград",
			"Кетчуп",
			"Знания",
			"Модель",
			"Минус",
			"Стол1",
			"Стул1",
			"Пирог1",
			"Волна1",
			"Арбуз1",
			"Дыня1",
			"Помидор1",
			"Игорь1",
			"Виноград1",
			"Кетчуп1",
			"Знания1",
			"Модель1",
			"Минус1",
		}

		for _, v := range title {
			//создание экземпляра книги
			book := model.ModelBook{
				Title:          v,
				Author:         "Шекспир",
				Published_year: 1956,
				Read:           false,
				Receiving_at:   time.Now(),
			}

			//вызов функции добавления нового экземпляра книги
			if err := interaction_db.Insert_book(ctx, conn, book); err != nil {
				log.Fatal(err)
			}
		}

		now := time.Now()
		bookNO := model.ModelBook{
			Title:          "Слон",
			Author:         "Шекспир",
			Published_year: 1956,
			Read:           true,
			Review: sql.NullString{
				String: "Готово",
				Valid:  true,
			},
			Receiving_at: time.Now(),
			Read_at:      &now,
			ID:           1,
		}

		if err := interaction_db.Edit_Book(ctx, conn, bookNO); err != nil {
			log.Fatal(err)
		}

		if err := interaction_db.Get_All_Books(ctx, conn); err != nil {
			log.Fatal(err)
		}

		booksIDs := []int{
			1, 2, 3,
		}
		if err := interaction_db.Delete_Books(ctx, conn, booksIDs); err != nil {
			log.Fatal(err)
		}
	*/

	if err := interaction_db.List_Pages(ctx, conn, 1); err != nil {
		log.Fatal(err)
	}

	log.Println("succeed")
}
