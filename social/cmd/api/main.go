package main

import "log"

func main() {
	cfg := config{
		addr: "127.0.0.1:8080",
	}
	app := application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
