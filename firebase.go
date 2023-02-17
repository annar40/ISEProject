package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Initialize the Firebase App
	config := &firebase.Config{
		DatabaseURL: "https://thoughtdump-4b31d-default-rtdb.firebaseio.com/",
	}
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, config, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Access the Firebase database
	db, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("error getting database: %v\n", err)
	}
	defer db.Client().Close()

	// Add data to the database
	ref := db.NewRef("/users")
	if err := ref.Set(ctx, map[string]string{"username": "johndoe"}); err != nil {
		log.Fatalf("error setting value: %v\n", err)
	}
}
