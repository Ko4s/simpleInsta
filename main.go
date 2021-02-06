package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var homeTemplate *template.Template

// w -> writer -> miejsce do którego wypiszemy wynik naszej funkcji
// wynikiem może być teskt, html, lub inny fomrat danych
// req -> struct opisuje rządanie do naszego programu

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeTemplate.Execute(w, nil)
}

func user(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "<h1>Hello, my user </h1>")
}

func main() {

	//Pasrowanie szablonów(templaty) html
	homeTemplate = template.Must(template.ParseFiles("./view/home.html"))

	// mux => HTTP request multiplexer => program który obsługuje  requesty od klientów
	// dopasowuje url do opowiedniej funkcji
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/contact/", user)

	fmt.Println("Server start with a bump")

	http.ListenAndServe("localhost:3000", mux)
}
