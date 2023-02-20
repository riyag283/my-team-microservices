package db

import (
	"log"
	util "teams/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DBClient *sqlx.DB

func InitialiseDBConnection() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations:", err)
	}

	db, err := sqlx.Open(config.DBDriver, config.DBSource)
	if err != nil {
		panic(err.Error())
	} 
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db 
}