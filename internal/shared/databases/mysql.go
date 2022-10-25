package databases

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConn() (*sql.DB, error) {
	connURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	)
	return sql.Open("mysql", connURL)
}

// WrapDBError translate database error
// into HTTP response code
func WrapDBError(err error) int {
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 404
		}

		return 500
	}
	return 0
}
