Detail work you've completed in Sprint 2

We had many troubles connecting our frontend and backend together. 

List unit tests and Cypress test for frontend

List unit tests for backend



**API Documentation:**

This code package implements a simple HTTP server that allows users to sign up and login. The server uses Firebase to store user data and Google Cloud's Firestore client to interact with Firebase. The server is designed to handle two API endpoints: /signup and /login.

**Signup Endpoint**
This endpoint allows users to sign up by sending a POST request with the user's information in JSON format.

Request
HTTP method: POST
Endpoint: /signup
Request body:

JSON: {
    "name": "string",
    "phone": "string",
    "email": "string",
    "password": "string"
}

**Response**
HTTP status code: 200 OK
Response body: "User data written to Firestore with ID: {ID}", where {ID} is the ID of the Firestore document that was created for the user.
If there is an error, the server will respond with an HTTP status code of 500 Internal Server Error.




**Login Endpoint**
This endpoint allows users to log in by sending a POST request with the user's name, email, and password in JSON format. The server will check if the email and password match the user's information in Firebase.

**Request**
HTTP method: POST
Endpoint: /login
Request body:

JSON:{
    "name": "string",
    "email": "string",
    "password": "string"
}

**Response**
HTTP status code: 200 OK
Response body: "Login successful"
If the email and password do not match the user's information in Firebase, the server will respond with an HTTP status code of 401 Unauthorized.
If there is an error, the server will respond with an HTTP status code of 500 Internal Server Error.
Firebase Initialization
The main function initializes Firebase with the project ID and a path to the service account key file. The Firestore client is then created with the initialized app.

**CORS**
The main function uses the rs/cors package to enable Cross-Origin Resource Sharing (CORS) for the HTTP server.

**HTTP Server**
The main function starts an HTTP server that listens on port 8000. The server handles requests to the /signup and /login endpoints using the signupHandler and loginHandler functions, respectively.

