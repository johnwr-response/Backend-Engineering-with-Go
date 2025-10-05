package main

import (
	"log"

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
	}
	store := internalstore.NewStorage(nil)
	app := application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
