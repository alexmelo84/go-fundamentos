package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, mundo!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listagem de usuários"))
}
