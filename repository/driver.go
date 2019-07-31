package repository

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"os"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

func init() {
	_ = gotenv.Load()
	pgUrl, err := pq.ParseURL(os.Getenv("POSTGRES_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	sourceUrl := os.Getenv("MIGRATION_SOURCE_URL")
	migration(sourceUrl)
}

func GetDB() *sql.DB {
	return db
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func migration(sourceUrl string) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	logFatal(err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+sourceUrl,
		"postgres", driver,
	)
	_ = m.Steps(1)
}
