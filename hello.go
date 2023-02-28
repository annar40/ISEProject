package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// User represents a user in the system
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", "root:anna123@tcp(localhost:3306)/users")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a HTTP server to handle requests
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body into a User struct
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Insert the new user into the database
		stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to MySQL database!")
		defer stmt.Close()

		result, err := stmt.Exec(user.Username, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		lastID, err := result.LastInsertId()
		// Return a success message
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rowsAffected != 1 {
			http.Error(w, "Error inserting user into database", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User  %s created successfully  with ID %d", user.Username, lastID)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
