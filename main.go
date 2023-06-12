package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello")
	})

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
