package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome string
	Email string
}

func main() {
	http.HandleFunc("/home", home)

	fmt.Println("Listening port 5000...")

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob("*.html"))

	u := usuario{"Alex Melo", "alex@iajsiajsia.com"}

	templates.ExecuteTemplate(w, "home.html", u)
}
