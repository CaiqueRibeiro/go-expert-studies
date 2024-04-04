package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books/{id}", GetBookHandler)
	// can capture everything that initiates with d, as /books/dir/djdjdjd/a/123 and uses as a single parameter
	mux.HandleFunc("GET /books/dir/{d...}", GetBookHandler)
	mux.HandleFunc("GET /books/{$}", BooksHandler) // ignore everyhing after /books/, gets exact path and gives 404 if other path is given

	http.ListenAndServe(":8080", mux)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Book ID: " + id))
}

func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	dirpath := r.PathValue("d") // Access captured directory path segments as slice
	fmt.Fprintf(w, "Accessing directory path %s\n", dirpath)
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books"))
}
