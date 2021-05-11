package dummy

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	mysqlM "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var db *sqlx.DB

func init() {
	db = createConnection()
	migrateDb()
}

func createConnection() *sqlx.DB {
	dbx, err := sqlx.Connect("mysql", "root:password@tcp(localhost:11000)/localdb?multiStatements=true&parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	dbx.SetConnMaxLifetime(time.Hour)
	dbx.SetMaxIdleConns(2)
	dbx.SetMaxOpenConns(5)

	return dbx
}

func migrateDb() {
	driver, err := mysqlM.WithInstance(db.DB, &mysqlM.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://resources/db", "localdb", driver)

	if err != nil {
		log.Fatalln(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalln(err)
	}
}

type Dummy struct {
	ID    int    `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

func FetchAllFromConnectionPool() ([]Dummy, error) {
	var dummies []Dummy
	err := db.Select(&dummies, "SELECT * FROM dummy")

	if err != nil {
		return nil, err
	}

	return dummies, nil
}

func FetchAllNewConnection() ([]Dummy, error) {
	db := createConnection()
	var dummies []Dummy
	err := db.Select(&dummies, "SELECT * FROM dummy")

	if err != nil {
		return nil, err
	}

	return dummies, nil
}
