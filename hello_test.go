package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"test/main"
	"testing"
)

func TestLogin(t *testing.T) {
	// create a request object for GET method
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// create a http.HandlerFunc for the handler function
	handler := http.HandlerFunc(Login)

	// serve the request using the handler function
	handler.ServeHTTP(rr, req)

	// check if the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// create a request object for POST method
	form := url.Values{}
	form.Add("username", "ashley")
	form.Add("password", "kelly")

	req, err = http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
