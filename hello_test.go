// package main

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	firebase "firebase.google.com/go"
// 	"google.golang.org/api/option"
// )

// func TestSignupHandler(t *testing.T) {
// 	// Initialize Firebase app
// 	ctx := context.Background()
// 	config := &firebase.Config{
// 		ProjectID: "thoughtdump-4b31d",
// 	}
// 	opt := option.WithCredentialsFile("C:/Users/arude/ThoughtDump/serviceAccountKey.json")
// 	app, err := firebase.NewApp(ctx, config, opt)
// 	if err != nil {
// 		t.Fatalf("error initializing app: %v", err)
// 	}

// 	// Create Firestore client
// 	client, err := app.Firestore(ctx)
// 	if err != nil {
// 		t.Fatalf("error creating Firestore client: %v", err)
// 	}
// 	defer client.Close()

// 	// Create a request body
// 	user := User{
// 		Name:     "testuser",
// 		Phone:    "1234567890",
// 		Email:    "test@example.com",
// 		Password: "testpassword",
// 	}
// 	body, err := json.Marshal(user)
// 	if err != nil {
// 		t.Fatalf("error marshaling user data: %v", err)
// 	}

// 	// Create a new HTTP request
// 	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatalf("error creating HTTP request: %v", err)
// 	}

// 	// Create a new HTTP recorder
// 	rec := httptest.NewRecorder()

// 	// Create a new HTTP handler and serve the request
// 	http.DefaultServeMux.ServeHTTP(rec, req)

// 	// Check the response status code
// 	if rec.Code != http.StatusOK {
// 		t.Errorf("unexpected response status code: got %d, want %d", rec.Code, http.StatusOK)
// 	}

// 	// Check the response body
// 	expectedResponse := "User data written to Firestore with ID:"
// 	if !bytes.Contains(rec.Body.Bytes(), []byte(expectedResponse)) {
// 		t.Errorf("unexpected response body: got %q, want to contain %q", rec.Body.String(), expectedResponse)
// 	}
// }

// // func TestLoginHandler(t *testing.T) {
// // 	// Initialize Firebase app
// // 	ctx := context.Background()
// // 	config := &firebase.Config{
// // 		ProjectID: "thoughtdump-4b31d",
// // 	}
// // 	opt := option.WithCredentialsFile("C:/Users/arude/ThoughtDump/serviceAccountKey.json")
// // 	app, err := firebase.NewApp(ctx, config, opt)
// // 	if err != nil {
// // 		t.Fatalf("error initializing app: %v", err)
// // 	}

// // 	// Create Firestore client
// // 	client, err := app.Firestore(ctx)
// // 	if err != nil {
// // 		t.Fatalf("error creating Firestore client: %v", err)
// // 	}
// // 	defer client.Close()

// // 	// Add a test user to the Firestore database
// // 	user := User{
// // 		Name:     "testuser",
// // 		Phone:    "1234567890",
// // 		Email:    "test@example.com",
// // 		Password: "testpassword",
// // 	}
// // 	docRef, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
// // 		"name":     user.Name,
// // 		"phone":    user.Phone,
// // 		"email":    user.Email,
// // 		"password": user.Password,
// // 	})
// // 	if err != nil {
// // 		t.Fatalf("error adding test user to Firestore: %v", err)
// // 	}

// // 	// Create a request body
// // 	loginUser := User{
// // 		Name:     "testuser",
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// func TestSignupHandler(t *testing.T) {
// 	// Create a test user
// 	user := User{
// 		Name:     "John Doe",
// 		Phone:    "555-555-5555",
// 		Email:    "johndoe@example.com",
// 		Password: "password123",
// 	}

// 	// Convert user to JSON
// 	userJSON, err := json.Marshal(user)
// 	if err != nil {
// 		t.Fatalf("error marshaling user data: %v", err)
// 	}

// 	// Create a new request with the user JSON data
// 	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(userJSON))
// 	if err != nil {
// 		t.Fatalf("error creating request: %v", err)
// 	}

// 	// Create a response recorder to record the response
// 	rr := httptest.NewRecorder()

// 	// Call the signup handler with the request and response recorder
// 	handler := http.HandlerFunc(main.signup)
// 	handler.ServeHTTP(rr, req)

// 	// Check the status code of the response
// 	if rr.Code != http.StatusOK {
// 		t.Errorf("expected status code %d, but got %d", http.StatusOK, rr.Code)
// 	}

//		// Check the response body
//		expectedBody := "User data written to Firestore with ID:"
//		if !bytes.Contains(rr.Body.Bytes(), []byte(expectedBody)) {
//			t.Errorf("expected response body to contain %q, but got %q", expectedBody, rr.Body.String())
//		}
//	}
// func TestSignupHandler(t *testing.T) {
// 	// Create a new HTTP request
// 	user := User{
// 		Name:     "John Doe",
// 		Phone:    "1234567890",
// 		Email:    "john.doe@example.com",
// 		Password: "password",
// 	}
// 	requestBody, err := json.Marshal(user)
// 	if err != nil {
// 		t.Fatalf("error marshaling request body: %v", err)
// 	}
// 	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		t.Fatalf("error creating HTTP request: %v", err)
// 	}

// 	// Send the HTTP request
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(main.signupHandler)
// 	handler.ServeHTTP(rr, req)

// 	// Check the HTTP response
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
// 	}
// 	expectedResponse := fmt.Sprintf("User data written to Firestore with ID: %v", "some_document_id")
// 	if rr.Body.String() != expectedResponse {
// 		t.Errorf("handler returned unexpected response: got %v, want %v", rr.Body.String(), expectedResponse)
// 	}
// }

func signupHandler(w http.ResponseWriter, r *http.Request) {
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

// func TestSignupHandler(t *testing.T) {
// 	// Create a new HTTP request
// 	user := User{
// 		Name:     "John Doe",
// 		Phone:    "1234567890",
// 		Email:    "john.doe@example.com",
// 		Password: "password",
// 	}
// 	requestBody, err := json.Marshal(user)
// 	if err != nil {
// 		t.Fatalf("error marshaling request body: %v", err)
// 	}
// 	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		t.Fatalf("error creating HTTP request: %v", err)
// 	}

// 	// Send the HTTP request
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(signupHandler)
// 	handler.ServeHTTP(rr, req)

//		// Check the HTTP response
//		if status := rr.Code; status != http.StatusOK {
//			t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
//		}
//		expectedResponse := fmt.Sprintf("User data written to Firestore with ID: %v", "some_document_id")
//		if rr.Body.String() != expectedResponse {
//			t.Errorf("handler returned unexpected response: got %v, want %v", rr.Body.String(), expectedResponse)
//		}
//	}
func TestSignupHandler(t *testing.T) {
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

	// Create a new request with a JSON body
	user := User{Name: "Test User", Phone: "555-1234", Email: "test@example.com", Password: "password"}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("error marshaling JSON: %v", err)
	}
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Call the SignupHandler function with the request and response recorder
	signupHandler(rr, req, client)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "User data written to Firestore with ID:"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
