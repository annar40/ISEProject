package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/rs/cors"
	"google.golang.org/api/option"
)

type User struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var ctx = context.Background()

func main() {
	// Initialize Firebase app
	ctx := context.Background()
	config := &firebase.Config{
		ProjectID: "thoughtdump-4b31d",
	}
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
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

	// Attach signup handler to HTTP server
	http.Handle("/signup", c.Handler(http.HandlerFunc(signupHandler(client))))

	// Attach login handler to HTTP server
	http.Handle("/login", c.Handler(http.HandlerFunc(loginHandler(client))))

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func signupHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse form data
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "error parsing form data", http.StatusBadRequest)
			return
		}

		// Write user data to Firestore
		_, err := client.Collection("users").Doc(user.Name).Set(ctx, map[string]interface{}{

			"name": user.Name,
			"email":    user.Email,
			"password": user.Password,
		})
		if err != nil {
			http.Error(w, "error writing user data to Firestore", http.StatusInternalServerError)
			return
		}

		// Send success response
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "User data written to Firestore")
	}
}
func loginHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, "Login Successful")
	}
}
