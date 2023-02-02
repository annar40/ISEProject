package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var userDB = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	case http.MethodPost:
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if p, ok := userDB[username]; ok && p == password {
			fmt.Fprintln(w, "Login Successful")
		} else {
			fmt.Fprintln(w, "Login Failed")
		}
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles("signup.html")
		t.Execute(w, nil)
	case http.MethodPost:
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if _, ok := userDB[username]; !ok {
			userDB[username] = password
			fmt.Fprintln(w, "Signup Successful")
		} else {
			fmt.Fprintln(w, "Username already taken")
		}
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8000", nil)
}
