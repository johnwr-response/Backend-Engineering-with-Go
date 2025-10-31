package main

import (
	"database/sql"
	"log"

	"github.com/johnwr-response/Backend-Engineering-with-Go/social/internal/db"
	"github.com/johnwr-response/Backend-Engineering-with-Go/social/internal/env"
	internalStore "github.com/johnwr-response/Backend-Engineering-with-Go/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminPassword@localhost/social_network?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	store := internalStore.NewStorage(conn)
	db.Seed(store)
}
