package main

import (
	// "context"
	// "fmt"
	// "html/template"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/gorilla/mux"
)

// var userDB = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
func (a *App) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	u.Username = uuid.New().String()
	err = a.db.Save(&u).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	// switch r.Method {
	// case http.MethodGet:
	// 	t, _ := template.ParseFiles("login.html")
	// 	t.Execute(w, nil)
	// case http.MethodPost:
	// 	r.ParseForm()
	// 	username := r.Form.Get("username")
	// 	password := r.Form.Get("password")

	// connect to firebase!
	// ctx := context.Background()
	// config := &firebase.Config{ProjectID: "thoughtdump-4b31d"}
	// sa := option.WithCredentialsFile("C:/Projects/ISEProject/ISEProject/serviceAccountKey.json")
	// app, err := firebase.NewApp(ctx, config, sa)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// 	return
	// }

	// client, err := app.Firestore(ctx)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// 	return
	// }

	// 	defer client.Close()

	// 	// Check if the user exists in Firebase
	// 	docRef := client.Doc("users/" + username)
	// 	snapshot, err := docRef.Get(ctx)
	// 	if err != nil {
	// 		fmt.Fprintln(w, "Login Failed- error in retrieving data")
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}

	// 	if snapshot == nil {
	// 		fmt.Fprintln(w, "Login Failed - document doesnt exist")
	// 		return
	// 	}

	// 	var user User
	// 	if err := snapshot.DataTo(&user); err != nil {
	// 		fmt.Fprintln(w, "Login Failed - snapshot data cant be converted to struct type")
	// 		return
	// 	}

	// 	if user.Password == password {
	// 		fmt.Fprintln(w, "Login Successful")
	// 	} else {
	// 		fmt.Fprintln(w, "Login Failed - password doesnt match")
	// 	}

	// }
	// }

	// func signup(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		t, _ := template.ParseFiles("app.component.html")
	// 		t.Execute(w, nil)
	// 	case http.MethodPost:
	// 		r.ParseForm()
	// 		username := r.Form.Get("username")
	// 		password := r.Form.Get("password")
	// 		if _, ok := userDB[username]; !ok {
	// 			userDB[username] = password
	// 			fmt.Fprintln(w, "Signup Successful")

	// 			// Connect to Firebase
	// 			ctx := context.Background()
	// 			config := &firebase.Config{ProjectID: "thoughtdump-4b31d"}
	// 			sa := option.WithCredentialsFile("C:/Projects/ISEProject/ISEProject/serviceAccountKey.json")
	// 			app, err := firebase.NewApp(ctx, config, sa)
	// 			if err != nil {
	// 				fmt.Printf("error: %v\n", err)
	// 				return
	// 			}

	// 			client, err := app.Firestore(ctx)
	// 			if err != nil {
	// 				fmt.Printf("error: %v\n", err)
	// 				return
	// 			}
	// 			defer client.Close()

	// 			// Write the new user to Firebase
	// 			docRef := client.Collection("users").Doc(username)
	// 			_, err = docRef.Set(ctx, map[string]interface{}{
	// 				"username": username,
	// 				"password": password,
	// 			})
	// 			if err != nil {
	// 				fmt.Printf("error: %v\n", err)
	// 				return
	// 			}
	// 		} else {
	// 			fmt.Fprintln(w, "Username already taken")
	// 		}
	// 	}
}

type App struct {
	db *gorm.DB
	r  *mux.Router
}

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/login", login).Methods("POST")
	// r.HandleFunc("/signup", signup).Methods("POST")
	// http.ListenAndServe(":8000", r)
	// http.HandleFunc("/login", login)
	// http.HandleFunc("/signup", signup)
	// http.ListenAndServe(":8000", nil)
	pass := os.Getenv("DB_PASS")
	db, err := gorm.Open(
		"postgres",
		"host=students-db user=go password="+pass+" dbname=go sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	app := App{
		db: db,
		r:  mux.NewRouter(),
	}
	db.AutoMigrate(&User{})
	app.r.HandleFunc("/login", app.login).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", app.r))

}
