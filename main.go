package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Home page!"))
}
func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from About page!"))
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Contact page!"))
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying specific snippet with id %d", id)
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
