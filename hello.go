package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
type Entry struct {
	JournalEntry string `json:"journalEntry"`
}

type Date struct {
	DateSelected string `json:"date"`
}

var currentUser string

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

	// Attach journal handler to HTTP server
	http.Handle("/journalEntry", c.Handler(http.HandlerFunc(journalHandler(client))))

	// Attach entry retriever handler to HTTP server
	http.Handle("/retrieveEntry", c.Handler(http.HandlerFunc(retrieveEntryHandler(client))))

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

			"name":     user.Name,
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
		currentUser = user.Name
	}
}

func journalHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse form data
		var entry Entry
		if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
			http.Error(w, "error parsing form data", http.StatusBadRequest)
			return
		}
		now := time.Now()
		dateStr := now.Format("2006-01-02") // Format the current date as "yyyy-mm-dd"

		// Write user data to Firestore
		_, err := client.Collection("users").Doc(currentUser).Collection("JournalEntry").Doc(dateStr).Set(ctx, map[string]interface{}{

			"journalEntry": entry.JournalEntry,
		})
		if err != nil {
			http.Error(w, "error writing user data to Firestore", http.StatusInternalServerError)
			return
		}
		if err != nil {
			http.Error(w, "error writing entry data to Firestore", http.StatusInternalServerError)
			return
		}

		// Send success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Entry data written to Firestore")
	}
}

func retrieveEntryHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var date Date
		if err := json.NewDecoder(r.Body).Decode(&date); err != nil {
			http.Error(w, "error parsing form data", http.StatusBadRequest)
			return
		}
		// Print the date
		fmt.Printf("Date selected: %v\n", date.DateSelected)

		// Get document with provided name
		docRef := client.Collection("users").Doc(currentUser).Collection("JournalEntry").Doc(date.DateSelected)

		// Get the data from the document
		docData, err := docRef.Get(ctx)
		if err != nil {
			log.Fatalf("Failed to get journal entry: %v", err)
		}

		// Get the "journalEntry" field from the document data
		journalEntry, exists := docData.Data()["journalEntry"]
		if !exists {
			log.Fatalf("Document does not have 'journalEntry' field")
		}

		// Print the journal entry
		fmt.Printf("Journal Entry: %s\n", journalEntry)

		// Send success response
		// w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "entry retrieved")
	}
}
