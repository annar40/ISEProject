package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"google.golang.org/api/iterator"

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

type Task struct {
	Text string `json:"text"`
}

const cookieSecretKey = "topSecretKey"

var ctx = context.Background()

// var store = sessions.NewCookieStore([]byte("secret-key"))

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

	// Attach task handler to HTTP server
	http.Handle("/entry", c.Handler(http.HandlerFunc(entryHandler(client))))

	//testing creating a cookie- WORKS!!

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// Create a new cookie
	// 	cookie := &http.Cookie{
	// 		Name:  "mycookie",
	// 		Value: "Hello World!",
	// 		Path:  "/",
	// 	}

	// 	// Set the cookie in the response header
	// 	http.SetCookie(w, cookie)

	// 	// Log the cookie value
	// 	// log.Printf("Cookie set: %v", cookie)

	// 	// Write a response
	// 	w.Write([]byte("Cookie set!"))
	// })

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
		userRef := client.Collection("users").Doc(user.Name)
		_, err := userRef.Set(ctx, map[string]interface{}{

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

		// Create a new SecureCookie instance
		// s := securecookie.New([]byte(cookieSecretKey), nil)

		// Encode the username as a cookie
		// encoded, err := s.Encode("username", user.Name)
		// if err != nil {
		// 	log.Fatalf("error creating cookie: %v", err)
		// }

		// Generate a unique session ID
		// sessionID := uuid.New().String()

		// Create a new cookie with the session ID as the value
		myCookie := &http.Cookie{
			Name:    "session",
			Value:   "temp",
			Path:    "/",
			Domain:  "localhost",
			Expires: time.Now().Add(24 * time.Hour),
		}

		// Set the cookie in the response header
		http.SetCookie(w, myCookie)

		cookies := r.Cookies()
		for _, cookie := range cookies {
			fmt.Println(cookie.Name, cookie.Value)
		}

		// Send success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Login Successful")
	}
}

func entryHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user name from signed cookie
		userName, err := getUsernameFromCookie(r)
		if err != nil {
			http.Error(w, "error getting user name from cookie", http.StatusUnauthorized)
			return
		}

		switch r.Method {
		case "POST":
			// Parse the task data from the request body
			var taskData map[string]string
			if err := json.NewDecoder(r.Body).Decode(&taskData); err != nil {
				http.Error(w, "error parsing task data", http.StatusBadRequest)
				return
			}

			// Create a new task document under the user's tasks collection
			taskRef, _, err := client.Collection("users").Doc(userName).Collection("tasks").Add(ctx, taskData)
			if err != nil {
				http.Error(w, "error creating task document", http.StatusInternalServerError)
				return
			}

			// Return the task ID to the client
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Created task with ID: %s", taskRef.ID)

		case "GET":
			// Get all task documents under the user's tasks collection
			tasksQuery := client.Collection("users").Doc(userName).Collection("tasks").Documents(ctx)

			// Iterate over the task documents and encode them as JSON in the response
			var tasks []map[string]interface{}
			for {
				docSnap, err := tasksQuery.Next()
				if err == iterator.Done {
					break
				} else if err != nil {
					http.Error(w, "error getting tasks from Firestore", http.StatusInternalServerError)
					return
				}

				var taskData map[string]interface{}
				if err := docSnap.DataTo(&taskData); err != nil {
					http.Error(w, "error parsing task data from Firestore", http.StatusInternalServerError)
					return
				}

				tasks = append(tasks, taskData)
			}

			// Write the tasks as JSON to the response
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(tasks); err != nil {
				http.Error(w, "error encoding tasks as JSON", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "unsupported HTTP method", http.StatusMethodNotAllowed)
			return
		}
	}
}

func getUsernameFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return "", err
	}

	value := make(map[string]string)

	s := securecookie.New([]byte(cookieSecretKey), nil)
	if err = s.Decode("session", cookie.Value, &value); err != nil {
		return "", err
	}

	username, ok := value["username"]
	if !ok {
		return "", errors.New("username not found in cookie")
	}

	return username, nil
}
