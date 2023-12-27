package main

import (
	"fmt"
	"webblog/internal/config"
	"webblog/internal/storage/postgres"
)

func main() {
	cfg := config.MustLoad()

	db, err := postgres.New(cfg.Storage_name)
	if err != nil {
		fmt.Printf("failde to init storage: %s", cfg.Storage_name)
		return
	}

	_ = db
}
