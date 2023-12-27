package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

//TODO: make all vars from env file or config

const (
	host     = "localhost"
	port     = 8089
	user     = "postgres"
	password = "vany2003"
	dbname   = "Articles"
)

type Storage struct {
	db *sql.DB
}

// TODO: make a fabric method for working using storage name

func New(storage_name string) (*Storage, error) {
	const op = "internal.storage.postgresql.New"

	//TODO: pass must be outside the code!!

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	fmt.Printf("Connection to %d port of %s data base is successfull\n", port, dbname)
	return &Storage{db: db}, nil
}

func (s *Storage) SaveArticle(author string, topic string, content string) error {
	const op = "internal.storage.postgresql.SaveArticle"

	stmt := "INSERT INTO articles(author, topic, content, date) VALUES ($1, $2, $3, $4)"
	curr := time.Now()

	_, err := s.db.Exec(stmt, author, topic, content, fmt.Sprint(curr.Date()))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	fmt.Println("content has been added")

	return nil
}

// TODO: если нет такого id, то ничего не делаем!!!

func (s *Storage) DeleteArticle(ID string) error {
	const op = "internal.storage.postgresql.DeleteArticle"

	stmt := "DELETE FROM articles WHERE ID=$1"

	_, err := s.db.Exec(stmt, ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	fmt.Println("content has been deleted")

	return nil
}
