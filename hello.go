package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
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

			//connect to firebase!

			ctx := context.Background()
			config := &firebase.Config{ProjectID: "thoughtdump-4b31d"}
			sa := option.WithCredentialsFile("C:/Users/arude/Visual Studio Code/ThoughtDump/ISEProject/serviceAccountKey.json")
			app, err := firebase.NewApp(ctx, config, sa)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}

			client, err := app.Firestore(ctx)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}

			defer client.Close()

			// Add new user to Firebase
			_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
				"username": username,
				"password": password,
			})
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}

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
