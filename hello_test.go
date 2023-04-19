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

// var fakeEntry = " Hay Fever was a comical play with a strong message. The actors throughout this play steadily used others to gain their own attention. The message this play left behind after many laughs was do not use others, plain and simple.The space was a well-constructed thrust stage however, the stage was not raised off the ground and the seats appeared to be mobile which lead to the belief that this is an Environmental Theater. As for the layout of the stage, there was a couch and a table in center stage for the first two acts, a door to the unseen outside front of the house upstage right, there were two large double doors upstage center leading to the backyard garden, and a raised floor upstage left containing a piano, a bookcase and a door leading to the home library. Stage left there was a half spiral staircase that lead up to the second floor which had a painting and lights that extended across to stage right. The second floor also contained the bedrooms for the house but the doors were not in sight. The detail was greater than expected: the hardwood floor throughout the first floor seemed to be genuine, the actual bookcase filled with real books not stage books, and the extremely large double doors with translucent glass and floral designs lead to the Garden, which featured an array of plants behind the doors. The entire space was well lit for the duration of the play; the stage took place inside the Bliss home so the bright lighting added a confortable feel to the inside of the house. The theatre was nearly sold out; it was hard to spot an empty seat, including the balcony and lower level seating. The crowd contained nearly an even amount of student audience members and senior citizens; there was few in the audience that appeared middle aged.The play was well cast, the entire Bliss family did a great job showing how overly dramatic and deceitful the family had become. Each actor portrayed their character accurately; there were no standouts that did match the character. Each visitor in the first act seemed to be feasible as far as the possibility of a relationship was concerned. All of the family members acted genuine. There was one actor who stood out as a perfect fit for his character. Taylor Rascher played Simon Bliss and played him perfectly. Raschers character Simon was a young man who was dramatic and romantic so much so it was almost to the point of ridiculousness. Rascher stole the show in the first act when he was romancing with Michelle Luz, who plays Myra Arundel. Rascher was displaying his love for Michelle Luz (Myra) in the most dramatic of fashions; he was playing the Blis game of pretending to be in complete love with someone and then a moment later change his mind. Rascher was proclaiming his love with elegant speeches and coddling up next to her and eventually kissing her. Rascher made this scene particularly hilarious because he was completely over the top in typical Bliss fashion. Of course, all of this was for naught because Luz ended up kissing Joe Hubbar character David Bliss in act two and also Rascher proclaimed his engagement to Caitlin Stagemolles character, Jackie Coryton. Rascher made his voice more shrill and audible to appear to be more dramatic and create more attention for himself. In the same way, Rascher used grander gestures with his arms (flailing, waving, etc.) to cause a more dramatic effect. The play was well interpreted by all whom were involved. Steven Wrentmore, the Director, kept the 1920s feel by dressing in all 1920s costumes and everyone spoke as if they were living at the time. Michelle Bisbee, the scene designer, made the inside of the home appear 192 because everything was grand. The Bliss home was grand with the spiral staircase, the very large backdoors, and the eloquent piano. The actorsmannerisms seemed like they were portraying a silent film. In older movies, actors seemed very dramatic and had flamboyant actions to prove so; the actors in Hay Fever shared the same feel for the dramatics. As far as Stephen Wrentmore directing goes, he did an excellent job. The scene when Chris Karl (Richard) and Caitlin Stegemoller (Jackie) enter and are left alone to make small talk with each other is the best pertaining to directing. The two actors used the entire stage in this scene and were very awkward with one another. This was Wrentmor doing because you could tell he had a vision for this scene in particular because it seemed very crisp and well rehearsed. The actors played it perfect with the excessively long pauses in their awkward small talk that the crowd was laughing through the entire scene. It appeared that Wrentmore instructed the actors, to keep their pauses longer than natural to heighten the awkward tension in the scene, which made it brilliant.The blocking throughout the play worked with the floor plan very well. All of the blocking worked seamlessly; the actors were never out of sight or in awkward positioning (ex. turned around, talking to someone behind them while face forward, etc.) even during the second act in the first scene while all eight actors were on set. One part of the set that stood out was the staircase, it is obviously very large, but the way Owen Virgin followed Megan Davis up the stairs was seamless. They both walked up the stairs with footsteps I unison, and Owen Virgin was so focused on every detail of Davis, it almost screamed out how much he was infatuated with her. The overall mood that was portrayed by the combination of lighting, sound, set, and costumes was very light and cheerful; at no point did the mood drop to something darker or saddening. This is common with many comedies because it becomes hard to laugh if the overall mood is down and dreary. The theatre space was very personal. First of all, it takes place in someone home so it is immediately personal. Also, the stage was built into the crowd just about so the audience felt like they were living the action out as it unfolded. The scenic design showed the audience without a doubt it was the 192s, with the barometer on the wall, the staircase, the piano, and the lights upstairs. However, there was little evidence to show what location the play took place."
var fakeMood = "😁 Great"

var fakeEntry = "Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test  Test test test test tes test test test DONE"

func TestLoginHandler(t *testing.T) {
	// Create a new request with a JSON body
	user := User{Name: "catherine", Email: "123@gmail.com", Password: "12345"}
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
	user := User{Name: "johnny", Email: "john@example.com", Password: "password123"}
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

func TestForDuplicateUsers(t *testing.T) {
	// Create a new request with a JSON body
	user := User{Name: "johnny", Email: "test@example.com", Password: "password"}
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
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusConflict)
	}
	expected := "username already taken"
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
}
func TestJournalHandler(t *testing.T) {

	// Create a new login request with a JSON body
	user := User{Name: "catherine", Email: "123@gmail.com", Password: "12345"}
	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(jsonBody))

	// Create a new recorder for capturing the response of login
	rr := httptest.NewRecorder()

	// Create a new journal entry request with a JSON body
	entry := Entry{JournalEntry: fakeEntry, Mood: fakeMood}
	jsonBody2, _ := json.Marshal(entry)
	req2, _ := http.NewRequest("POST", "/journalEntry", bytes.NewReader(jsonBody2))

	// Create a new recorder for capturing the response of entry
	rr2 := httptest.NewRecorder()

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
	loginhandler := http.HandlerFunc(loginHandler(client))
	c := cors.Default().Handler(loginhandler)
	c.ServeHTTP(rr, req)

	// Call the journal handler
	entryhandler := http.HandlerFunc(journalHandler(client))
	c2 := cors.Default().Handler(entryhandler)
	c2.ServeHTTP(rr2, req2)

	// Check that the response is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	expected := "Login Successful"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
	// Check that the response is as expected
	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	expected2 := "Entry data written to Firestore"
	if rr2.Body.String() != expected2 {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr2.Body.String(), expected2)
	}
}

func TestRetrieveEntryHandler(t *testing.T) {

	// Create a new login request with a JSON body
	user := User{Name: "catherine", Email: "123@gmail.com", Password: "12345"}
	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(jsonBody))

	// Create a new recorder for capturing the response of login
	rr := httptest.NewRecorder()

	// Create a new journal entry request with a JSON body
	entry := Entry{JournalEntry: fakeEntry, Mood: fakeMood}
	jsonBody2, _ := json.Marshal(entry)
	req2, _ := http.NewRequest("POST", "/journalEntry", bytes.NewReader(jsonBody2))

	// Create a new recorder for capturing the response of entry
	rr2 := httptest.NewRecorder()

	// Create a new journal entry request with a JSON body
	dateSelected := Date{DateSelected: "2023-03-29"}
	jsonBody3, _ := json.Marshal(dateSelected)
	req3, _ := http.NewRequest("POST", "/retrieveEntry", bytes.NewReader(jsonBody3))

	// Create a new recorder for capturing the response of entry
	rr3 := httptest.NewRecorder()

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
	loginhandler := http.HandlerFunc(loginHandler(client))
	c := cors.Default().Handler(loginhandler)
	c.ServeHTTP(rr, req)

	// Call the journal handler
	entryhandler := http.HandlerFunc(journalHandler(client))
	c2 := cors.Default().Handler(entryhandler)
	c2.ServeHTTP(rr2, req2)

	// Call the journal handler
	retrieveEntryhandler := http.HandlerFunc(retrieveEntryHandler(client))
	c3 := cors.Default().Handler(retrieveEntryhandler)
	c3.ServeHTTP(rr3, req3)

	// Check that the response is as expected
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	expected := "Login Successful"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
	// Check that the response is as expected
	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	expected2 := "Entry data written to Firestore"
	if rr2.Body.String() != expected2 {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr2.Body.String(), expected2)
	}

	// Check that the response is as expected
	if status := rr3.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected3 := "{\"dateSelected\":\"2023-03-29\",\"entry\":\"Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test Test test test test tes test test test test test test test test test test test test  Test test test test tes test test test DONE\",\"mood\":\"😁 Great\"}"
	if rr3.Body.String() != expected3 {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr3.Body.String(), expected3)
	}

	//clicking a date that does not have an entry should return null
	dateSelected = Date{DateSelected: "2023-01-29"}
	// Call the journal handler
	retrieveEntryhandler = http.HandlerFunc(retrieveEntryHandler(client))
	c3 = cors.Default().Handler(retrieveEntryhandler)
	c3.ServeHTTP(rr3, req3)
	// Create a new recorder for capturing the response of entry
	rr3 = httptest.NewRecorder()
	expected4 := ""
	if rr3.Body.String() != expected4 {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr3.Body.String(), expected4)
	}
}

func TestRetrieveMoodsHandler(t *testing.T) {

	// Create a new request with a JSON body
	user := User{Name: "austin", Email: "a@a", Password: "12345"}
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

	// Create HTTP request and response objects
	req, err = http.NewRequest("GET", "/retrieve-moods", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr = httptest.NewRecorder()

	// Call the handler function
	handler = retrieveMoodsHandler(client)
	handler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `{"moods":{"😡 Angry":1}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)

	}
}
