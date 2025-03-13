package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (a *api) handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from " + r.URL.Path))
	log.Default().Printf("%v %v", r.Method, r.URL.Path)
}

func (a *api) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from get users"))
	log.Default().Printf("%v %v", r.Method, r.URL.Path)
}

func main() {
	api := &api{addr: ":8080"}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", api.handleRequest)
	mux.HandleFunc("POST /users", api.handleGetUsers)
	mux.HandleFunc("GET /users", api.handleGetUsers)

	srv := http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	log.Printf("Listening to port %v\n", api.addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
