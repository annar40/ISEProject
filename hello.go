package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// User struct to hold login credentials
type User struct {
	UserName string
	Password string
}

// login page handler function
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// extract the username and password from the form
		u := User{
			UserName: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		// check if the login credentials are correct
		if u.UserName == "admin" && u.Password == "password" {
			fmt.Fprint(w, "Welcome, "+u.UserName+"! You are logged in.")
		} else {
			fmt.Fprint(w, "Invalid login credentials. Please try again.")
		}
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8000", nil)
}
