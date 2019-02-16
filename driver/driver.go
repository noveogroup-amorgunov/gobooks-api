package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"gobooks-api/utils"
	"os"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("POSTGRES_URI"))
	utils.LogFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	utils.LogFatal(err)

	err = db.Ping()
	utils.LogFatal(err)

	return db;
}
