package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	firebase "firebase.google.com/go"
	"github.com/rs/cors"

	// "thoughtDump/go/pkg/mod/cloud.google.com/go/firestore@v1.6.1"

	"google.golang.org/api/option"
)

func TestLoginHandler(t *testing.T) {
	// Create a new request with a JSON body
	user := User{Name: "catherine", Email: "gearshift2021@gmail.com", Password: "12345"}
	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(jsonBody))

	// Create a new recorder for capturing the response
	rr := httptest.NewRecorder()

	// Initialize Firebase app
	ctx := context.Background()
	config := &firebase.Config{
		ProjectID: "thoughtdump-4b31d",
	}
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		t.Fatalf("error initializing app: %v", err)
	}

	// Create Firestore client
	client, err := app.Firestore(ctx)
	if err != nil {
		t.Fatalf("error creating Firestore client: %v", err)
	}
	defer client.Close()

	// Call the login handler
	handler := http.HandlerFunc(loginHandler(client))
	c := cors.Default().Handler(handler)
	c.ServeHTTP(rr, req)

	// Check that the response is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	expected := "Login Successful"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}

func TestSignupHandler(t *testing.T) {
	// Create a new request with a JSON body
	user := User{Name: "john", Email: "john@example.com", Password: "password123"}
	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewReader(jsonBody))

	// Create a new recorder for capturing the response
	rr := httptest.NewRecorder()

	// Initialize Firebase app and create Firestore client
	ctx := context.Background()
	config := &firebase.Config{ProjectID: "thoughtdump-4b31d"}
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		t.Fatalf("error initializing app: %v", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		t.Fatalf("error creating Firestore client: %v", err)
	}
	defer client.Close()

	// Call the signup handler
	handler := http.HandlerFunc(signupHandler(client))
	c := cors.Default().Handler(handler)
	c.ServeHTTP(rr, req)

	// Check that the response is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	expected := "User data written to Firestore"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}

	// Check that the user was actually added to the database
	docRef := client.Collection("users").Doc(user.Name)
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		t.Fatalf("error getting user data from Firestore: %v", err)
	}
	var userData User
	if err := docSnap.DataTo(&userData); err != nil {
		t.Fatalf("error parsing user data from Firestore: %v", err)
	}
	if userData.Email != user.Email || userData.Password != user.Password {
		t.Errorf("user not added to Firestore correctly: got %v, want %v", userData, user)
	}
}
