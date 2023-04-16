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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Entry struct {
	JournalEntry string `json:"text"`
	Mood         string `json:"mood"`
}

type Date struct {
	DateSelected string `json:"date"`
}
type JournalEntry struct {
	DateSelected string `json:"dateSelected"`
	Entry        string `json:"entry"`
	Mood         string `json:"mood"`
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
		// Check if username is available
		docRef := client.Collection("users").Doc(user.Name)
		_, err := docRef.Get(ctx)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				// Username is available
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
				return
			} else {
				http.Error(w, "error checking username availability", http.StatusInternalServerError)
				return
			}
		} else {
			// Username is already taken
			http.Error(w, "username already taken", http.StatusConflict)
			return
		}
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
			"mood":         entry.Mood,
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

		// Get document with provided name
		docRef := client.Collection("users").Doc(currentUser).Collection("JournalEntry").Doc(date.DateSelected)
		// Get the data from the document
		docData, err := docRef.Get(ctx)
		if status.Code(err) != codes.NotFound {
			if err != nil {
				log.Fatalf("Failed to get journal entry: %v", err)
			}

			// Get the "journalEntry" field from the document data
			journalEntry, exists := docData.Data()["journalEntry"]
			if !exists {
				log.Fatalf("Document does not have 'journalEntry' field")
			}

			moodEntry, exists := docData.Data()["mood"]
			if !exists {
				log.Fatalf("Document does not have 'journalEntry' field")
			}

			// Create JournalEntry struct
			entry := JournalEntry{
				DateSelected: date.DateSelected,
				Entry:        journalEntry.(string),
				Mood:         moodEntry.(string),
			}

			// Marshal JournalEntry struct as JSON
			jsonResponse, err := json.Marshal(entry)
			if err != nil {
				log.Fatalf("Failed to marshal JSON: %v", err)
			}

			// Send success response with JSON data
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
		}
	}
}
