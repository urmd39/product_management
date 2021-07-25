package control

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connected() *sql.DB {
	infoDB := "host=localhost user=postgres dbname=product_management password=model.kn2412" +
		" sslmode=disable"
	db, err := sql.Open("postgres", infoDB)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

// func Disconnect() {

// }
