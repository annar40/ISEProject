package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	opt := option.WithCredentialsFile("C:/Users/arude/ThoughtDump/serviceAccountKey.json")
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
	expected := "User data written to Firestore with ID"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}
