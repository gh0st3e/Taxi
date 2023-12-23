package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const dbLifetime = 20 * time.Second

func MySqlConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "a97166_taxiexp:Denis_2003_db@tcp(188.225.76.205)/db_a97166_taxiexp")
	if err != nil {
		return nil, fmt.Errorf("connect to DB error: %s", err)
	}
	db.SetConnMaxLifetime(dbLifetime)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping DB error: %s", err)
	}

	return db, nil
}
