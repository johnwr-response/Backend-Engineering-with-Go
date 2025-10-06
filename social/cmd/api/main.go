package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	internaldb "github.com/johnwr-response/Backend-Engineering-with-Go/social/internal/db"
	"github.com/johnwr-response/Backend-Engineering-with-Go/social/internal/env"
	internalstore "github.com/johnwr-response/Backend-Engineering-with-Go/social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("SERVER_ADDR", "8080"),
		db: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgres://admin:adminPassword@localhost/social_network?sslmode=disable"),
			maxOpenCons: env.GetInt("DB_MAX_OPEN_CONS", 30),
			maxIdleCons: env.GetInt("DB_MAX_IDLE_CONS", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := internaldb.New(
		cfg.db.addr,
		cfg.db.maxOpenCons,
		cfg.db.maxIdleCons,
		cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}(db)
	log.Println("database connection pool established")

	store := internalstore.NewStorage(db)
	app := application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
