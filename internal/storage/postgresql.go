package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 8089
	user     = "postgres"
	password = "vany2003"
	dbname   = "Logon"
)

type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	const op = "internal.storage.postgresql.New"

	//TODO: pass must be outside the code!!

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	fmt.Printf("Connection to %d port of %s data base is successfull", port, dbname)
	return &Storage{db: db}, nil
}
