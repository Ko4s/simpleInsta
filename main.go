package main

import (
	"fmt"
	"net/http"
)

// w -> writer -> miejsce do którego wypiszemy wynik naszej funkcji
// wynikiem może być teskt, html, lub inny fomrat danych
// req -> struct opisuje rządanie do naszego programu

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Welcome to my home page") // Fprintf wypisuje podany tekst na podane wyjscie
}

func user(w http.ResponseWriter, req *http.Request) {
	
	fmt.Fprintf(w, "<h1>Hello, my user </h1>")
}

func main() {
	// mux => HTTP request multiplexer => program który obsługuje  requesty od klientów
	// dopasowuje url do opowiedniej funkcji
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/user/", user)

	http.ListenAndServe("localhost:3000", mux)
	fmt.Println("Server start with a bump")

}
