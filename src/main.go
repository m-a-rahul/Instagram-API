package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Welcome to the Instagram API Backend application"}`))

}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users/", userRequestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}