Detail work you've completed in Sprint 2

We had many troubles connecting our frontend and backend together. 

List unit tests and Cypress test for frontend

List unit tests for backend



**API Documentation:**
Package Name: Firebase Auth API
Package Description: This package provides APIs to handle user authentication using Firebase and Firestore.

**Imported Packages:***

context: This package is used for passing context between functions.
encoding/json: This package is used to encode and decode JSON data.
fmt: This package is used for formatting output.
log: This package is used for logging error messages.
net/http: This package is used to implement HTTP client and server.
cloud.google.com/go/firestore: This package is used to interact with Firestore.
firebase.google.com/go: This package is used to initialize Firebase app.
github.com/rs/cors: This package is used to handle Cross-Origin Resource Sharing (CORS).
google.golang.org/api/option: This package is used to provide options for authentication and other settings.

**Structures:**

User: This structure is used to represent user information. It contains four fields:
Name: string, the name of the user.
Phone: string, the phone number of the user.
Email: string, the email address of the user.
Password: string, the password of the user.

**Variables:**

ctx: context.Context, a context for passing context between functions.

**Functions:**

main(): This function is the entry point of the program. It initializes the Firebase app and creates a Firestore client. Then it attaches signup and login handlers to the HTTP server and starts the server.

signupHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request): This function is used to handle user signup requests. It takes a Firestore client as an input parameter and returns a function that handles HTTP requests. This function parses form data, writes user data to Firestore, and sends a success response.

loginHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request): This function is used to handle user login requests. It takes a Firestore client as an input parameter and returns a function that handles HTTP requests. This function parses form data, retrieves user data from Firestore, checks if email and password match, and sends a success response if the login is successful.

**API Endpoints:**

/signup: This endpoint is used for user signup requests. It accepts a POST request with user data in JSON format. The user data must contain the name, email, and password fields. If the signup is successful, it sends a 200 OK response with the message "User data written to Firestore".

/login: This endpoint is used for user login requests. It accepts a POST request with user data in JSON format. The user data must contain the name, email, and password fields. If the login is successful, it sends a 200 OK response with the message "Login Successful".

**Sample Usage:**
To use this package, you can import it in your Go program and call the main function. You can then send HTTP requests to the /signup and /login endpoints to handle user signup and login requests.
