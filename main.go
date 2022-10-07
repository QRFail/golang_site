package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func initHadlers() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/about/", about_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	initHadlers()
}
