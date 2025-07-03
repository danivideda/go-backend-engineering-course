package main

import (
	"log"

	"github.com/danivideda/go-backend-engineering-course/social/internal/env"
	"github.com/danivideda/go-backend-engineering-course/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
