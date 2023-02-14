package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DBClient *sqlx.DB

func InitialiseDBConnection() {
	db, err := sqlx.Open("postgres", "postgres://postgres:password@localhost:5433/users?sslmode=disable")
	if err != nil {
		panic(err.Error())
	} 
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db 
}