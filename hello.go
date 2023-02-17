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

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

		//connect to firebase!
		ctx := context.Background()
		config := &firebase.Config{ProjectID: "thoughtdump-4b31d"}
		sa := option.WithCredentialsFile("C:/Projects/ISEProject/ISEProject/serviceAccountKey.json")
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

		// Check if the user exists in Firebase
		docRef := client.Doc("users/" + username)
		snapshot, err := docRef.Get(ctx)
		if err != nil {
			fmt.Fprintln(w, "Login Failed- error in retrieving data")
			fmt.Println("Error:", err)
			return
		}

		if snapshot == nil {
			fmt.Fprintln(w, "Login Failed - document doesnt exist")
			return
		}

		var user User
		if err := snapshot.DataTo(&user); err != nil {
			fmt.Fprintln(w, "Login Failed - snapshot data cant be converted to struct type")
			return
		}

		if user.Password == password {
			fmt.Fprintln(w, "Login Successful")
		} else {
			fmt.Fprintln(w, "Login Failed - password doesnt match")
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

			// Connect to Firebase
			ctx := context.Background()
			config := &firebase.Config{ProjectID: "thoughtdump-4b31d"}
			sa := option.WithCredentialsFile("C:/Projects/ISEProject/ISEProject/serviceAccountKey.json")
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

			// Write the new user to Firebase
			docRef := client.Collection("users").Doc(username)
			_, err = docRef.Set(ctx, map[string]interface{}{
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
