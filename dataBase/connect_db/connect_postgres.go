package connect_db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

//"postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

// подключение к базе данных
func Connect_db(ctx context.Context) (*pgx.Conn, error) {

	//возвращаю подключение
	return pgx.Connect(ctx, "postgres://postgres:2906@localhost:5432/postgres")
}
