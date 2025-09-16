package go_database_mysql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabase(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.Close()
}
