package main

import (
	"log"
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	srv := http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	log.Printf("Listening to port %v\n", api.addr)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
