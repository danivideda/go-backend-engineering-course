package main

import (
	"log"

	"github.com/danivideda/go-backend-engineering-course/social/internal/db"
	"github.com/danivideda/go-backend-engineering-course/social/internal/env"
	"github.com/danivideda/go-backend-engineering-course/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")

	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Panicln("Database conn failed:", err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store)
}
