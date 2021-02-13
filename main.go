package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var homeTemplate *template.Template
var contactTemplate *template.Template
var faqTemplate *template.Template

// w -> writer -> miejsce do którego wypiszemy wynik naszej funkcji
// wynikiem może być teskt, html, lub inny fomrat danych
// req -> struct opisuje rządanie do naszego programu

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	dumbUser := User{
		Name:  "Piotrek",
		Color: "Red",
		Age:   253623,
	}

	homeTemplate.ExecuteTemplate(w, "bootstrap", dumbUser)
}

func contact(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	contactTemplate.Execute(w, nil)
}

func faq(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	faqTemplate.Execute(w, nil)
}

type User struct {
	Name  string
	Color string
	Age   int
}

func main() {

	//Pasrowanie szablonów(templaty) html
	homeTemplate = template.Must(template.ParseFiles("./view/layouts/bootstrap.html", "./view/home.html"))
	contactTemplate = template.Must(template.ParseFiles("./view/contact.html"))
	faqTemplate = template.Must(template.ParseFiles("./view/faq.html"))

	// mux => HTTP request multiplexer => program który obsługuje  requesty od klientów
	// dopasowuje url do opowiedniej funkcji
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/contact/", contact)
	mux.HandleFunc("/faq/", faq)

	fmt.Println("Server start with a bump")

	http.ListenAndServe("localhost:3000", mux)
}
