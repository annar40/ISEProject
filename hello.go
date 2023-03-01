package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type User struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	// Initialize Firebase app
	ctx := context.Background()
	config := &firebase.Config{
		ProjectID: "thoughtdump-4b31d",
	}
	opt := option.WithCredentialsFile("C:/Users/arude/ThoughtDump/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	// Create Firestore client
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error creating Firestore client: %v", err)
	}
	defer client.Close()
	c := cors.Default()

	// Signup handler
	// http.Handle("/signup", c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	signupHandler := func(w http.ResponseWriter, r *http.Request) {
		// Parse form data
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "error parsing form data", http.StatusBadRequest)
			return
		}

		// Write user data to Firestore
		docRef, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
			"name":     user.Name,
			"phone":    user.Phone,
			"email":    user.Email,
			"password": user.Password,
		})
		if err != nil {
			http.Error(w, "error writing user data to Firestore", http.StatusInternalServerError)
			return
		}

		// Send success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User data written to Firestore with ID: %v", docRef.ID)
	}
	// Attach signup handler to HTTP server
	http.Handle("/signup", c.Handler(http.HandlerFunc(signupHandler)))

	// Login handler
	// http.Handle("/login", c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	loginHandler := func(w http.ResponseWriter, r *http.Request) {

		// Parse form data
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "error parsing form data", http.StatusBadRequest)
			return
		}

		// Get document with provided name
		docRef := client.Collection("users").Doc(user.Name)
		docSnap, err := docRef.Get(ctx)
		if err != nil {
			http.Error(w, "error getting user data from Firestore", http.StatusInternalServerError)
			return
		}

		// Check if email and password match
		var userData User
		if err := docSnap.DataTo(&userData); err != nil {
			http.Error(w, "error parsing user data from Firestore", http.StatusInternalServerError)
			return
		}
		if userData.Email != user.Email || userData.Password != user.Password {
			http.Error(w, "incorrect email or password", http.StatusUnauthorized)
			return
		}

		// Send success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Login successful")
	}
	// Attach signup handler to HTTP server
	http.Handle("/login", c.Handler(http.HandlerFunc(loginHandler)))
	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
