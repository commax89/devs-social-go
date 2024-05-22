package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("views/index.html"))

func handleIndex(w http.ResponseWriter, r *http.Request) {
	_ = templates.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	http.HandleFunc("/", handleIndex)
	fmt.Println("Server started on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
