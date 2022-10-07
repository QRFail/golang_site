package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name     string
	age      uint16
	happines float32
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("My name is %s. Happines: %f", u.name, u.happines)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, 0.15}

	bob.setNewName("Sam")

	fmt.Fprintf(w, bob.getAllInfo())
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
