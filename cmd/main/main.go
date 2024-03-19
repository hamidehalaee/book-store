package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hamidehalaee/go-bookStore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

// how to run : CGO_ENABLED=0 go run main.go
// how to run : CGO_ENABLED=0 go build main.go
