package db

import (
	"database/sql"
	"log"
	"os"
	util "teams/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBClientInterface interface {
    Ping() error
    Close() error
    Exec(query string, args ...interface{}) (sql.Result, error)
    Get(dest interface{}, query string, args ...interface{}) error
    Select(dest interface{}, query string, args ...interface{}) error
    QueryRow(query string, args ...interface{}) *sql.Row
    QueryRowx(query string, args ...interface{}) *sqlx.Row
    Query(query string, args ...interface{}) (*sql.Rows, error)
}

type DBClient struct {
    *sqlx.DB
}

func (c *DBClient) QueryRow(query string, args ...interface{}) *sql.Row {
    return c.DB.QueryRow(query, args...)
}

func (c *DBClient) QueryRowx(query string, args ...interface{}) *sqlx.Row {
    return c.DB.QueryRowx(query, args...)
}

func (c *DBClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
    return c.DB.Query(query, args...)
}

var (
    DBClientInstance DBClientInterface
)

func InitialiseDBConnection() {
    var config util.Config
    var err error

    if os.Getenv("APP_ENV") == "test" {
		config, err = util.LoadMockConfig()
	} else {
		config, err = util.LoadConfig(".")
	}

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

    DBClientInstance = &DBClient{db}
}
