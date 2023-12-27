package main

import (
	"fmt"
	"webblog/internal/config"
	"webblog/internal/storage"
)

func main() {
	// TODO: make a database PostgreSQL
	// TODO:
	cfg := config.MustLoad()

	fmt.Println(cfg)

	db, _ := storage.New()
	_ = db
}
