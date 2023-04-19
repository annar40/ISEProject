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
	Streak   int    `json:"streak"`
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

type EntryDate struct {
	Date string `json:"date"`
}

type Streak struct {
	StreakNum int64 `json:"streakNum"`
}

var currentUser string
var hasAStreak bool

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

	http.Handle("/retrieveDates", c.Handler(http.HandlerFunc(retrieveDatesHandler(client))))

	http.Handle("/retrieveMoods", c.Handler(http.HandlerFunc(retrieveMoodsHandler(client))))

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
					"streak":   0,
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
		hasAStreak = false
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

		// Write user data and update streak to Firestore
		docRef := client.Collection("users").Doc(currentUser)
		_, err := docRef.Collection("JournalEntry").Doc(dateStr).Set(ctx, map[string]interface{}{
			"journalEntry": entry.JournalEntry,
			"mood":         entry.Mood,
		})
		if err != nil {
			http.Error(w, "error writing entry data to Firestore", http.StatusInternalServerError)
			return
		}
		if hasAStreak {
			_, err = docRef.Update(ctx, []firestore.Update{{Path: "streak", Value: firestore.Increment(1)}})
		} else {

			_, err = docRef.Update(ctx, []firestore.Update{{Path: "streak", Value: 1}})
		}

		if err != nil {
			http.Error(w, "error updating streak field", http.StatusInternalServerError)
			return
		}

		userDoc, err := client.Collection("users").Doc(currentUser).Get(ctx)
		if err != nil {
			http.Error(w, "error retrieving user data", http.StatusInternalServerError)
			return
		}

		currentStreak := Streak{
			StreakNum: userDoc.Data()["streak"].(int64),
		}

		jsonResponse, err := json.Marshal(currentStreak)

		if err != nil {
			http.Error(w, "error marshaling dates into JSON", http.StatusInternalServerError)
			return
		}

		// Write JSON string to response body
		w.Header().Set("Content-Type", "application/json")
		// Send success response
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)

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

func retrieveDatesHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query Firestore to retrieve all journal entry documents of the current user
		docs, err := client.Collection("users").Doc(currentUser).Collection("JournalEntry").Documents(ctx).GetAll()
		if err != nil {
			http.Error(w, "error retrieving journal entries", http.StatusInternalServerError)
			return
		}

		// Get yesterday's date
		yesterday := getYesterday()

		// Check if the last entry in the array of dates is equal to yesterday's date
		lastEntry := len(docs) - 1

		userDoc, err := client.Collection("users").Doc(currentUser).Get(ctx)
		if err != nil {
			http.Error(w, "error retrieving user data", http.StatusInternalServerError)
			return
		}

		var currentStreak int64
		currentStreak = 0

		if lastEntry >= 0 {
			lastEntryDate := docs[lastEntry].Ref.ID
			if lastEntryDate == yesterday {
				hasAStreak = true
				// Get streak field from user document
				currentStreak = userDoc.Data()["streak"].(int64)
			}
		}

		// Extract IDs of documents, which correspond to dates of journal entries
		var dates []EntryDate
		for _, doc := range docs {
			dates = append(dates, EntryDate{Date: doc.Ref.ID})
		}

		// Marshal dates and isYesterdayEntry into a JSON string
		jsonBytes, err := json.Marshal(struct {
			Dates         []EntryDate `json:"dates"`
			CurrentStreak int64       `json:"CurrentStreak"`
		}{
			Dates:         dates,
			CurrentStreak: currentStreak,
		})
		if err != nil {
			http.Error(w, "error marshaling dates into JSON", http.StatusInternalServerError)
			return
		}
		jsonString := string(jsonBytes)

		// Write JSON string to response body
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsonString))
	}
}

func getYesterday() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	yesterdayString := yesterday.Format("2006-01-02")
	return yesterdayString
}

func retrieveMoodsHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query Firestore to retrieve all journal entry documents of the current user
		docs, err := client.Collection("users").Doc(currentUser).Collection("JournalEntry").Documents(ctx).GetAll()
		if err != nil {
			http.Error(w, "error retrieving journal entries", http.StatusInternalServerError)
			return
		}

		// Extract IDs of documents, which correspond to dates of journal entries
		var dates []EntryDate
		for _, doc := range docs {
			dates = append(dates, EntryDate{Date: doc.Ref.ID})
		}

		// Loop through dates and print moods =========================apply better logic in here
		for _, date := range dates {

			// Get document with provided name
			docRef := client.Collection("users").Doc(currentUser).Collection("JournalEntry").Doc(date.Date)
			// Get the data from the document
			docData, err := docRef.Get(ctx)
			if status.Code(err) != codes.NotFound {
				if err != nil {
					log.Fatalf("Failed to get journal entry: %v", err)
				}

				moodEntry, exists := docData.Data()["mood"]
				if !exists {
					log.Fatalf("Document does not have 'mood' field")
				}
				fmt.Println(moodEntry)
			}
		}

		// Marshal dates and isYesterdayEntry into a JSON string ========================change this completely
		jsonBytes, err := json.Marshal(struct {
			Dates []EntryDate `json:"dates"`
		}{
			Dates: dates,
		})
		if err != nil {
			http.Error(w, "error marshaling dates into JSON", http.StatusInternalServerError)
			return
		}
		jsonString := string(jsonBytes)

		// Write JSON string to response body
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsonString))
	}
}
