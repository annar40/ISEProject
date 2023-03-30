
## Video
link goes here


## Work Completed in Sprint 3
-store journal entries (track logged in user)

-retrieve journal entries

-specific changes to journal entry page

-rerouting from journal to entry

-new history page 

-fix unique user bug

-new about page

-new home page

-store user mood


## Frontend unit tests

## Backend unit tests
func TestJournalHandler(t *testing.T) : logs into existing account. Writes an entry. Stores entry under that user. 

func TestEntryRetrieverHandler() :blah blah to be written

## **Updated API Documentation:**

1.         /signup - POST
This endpoint allows a user to sign up. The request body should contain a JSON object with the following fields:
{
  "name": "string",
  "email": "string",
  "password": "string"
}
The name field is used as the document ID in Firestore.

On success, the endpoint returns a status code of 200 and the message "User data written to Firestore". On failure, the endpoint returns a status code of 400 for malformed request or 500 for internal server error.

2.       /login - POST
This endpoint allows a user to login. The request body should contain a JSON object with the following fields:
{
  "name": "string",
  "email": "string",
  "password": "string"
}

The name field is used as the document ID in Firestore.

On success, the endpoint returns a status code of 200 and the message "Login Successful". On failure, the endpoint returns a status code of 400 for malformed request or 401 for unauthorized access or 500 for internal server error.

3.      /journalEntry - POST
This endpoint allows a user to create a journal entry. The request body should contain a JSON object with the following fields:
{
  "text": "string",
  "mood": "string"
}
The text field is the journal entry text and the mood field is the mood for that entry.

On success, the endpoint returns a status code of 200 and the message "Entry data written to Firestore". On failure, the endpoint returns a status code of 400 for malformed request or 500 for internal server error.

4.     /retrieveEntry - POST
This endpoint allows a user to retrieve a journal entry for a specific date. The request body should contain a JSON object with the following field:
{
  "date": "string"
}
The date field should be in the format "yyyy-mm-dd".

On success, the endpoint returns a status code of 200 and a JSON object with the following fields:
{
  "dateSelected": "string",
  "entry": "string",
  "mood": "string"
}
On failure, the endpoint returns a status code of 400 for malformed request or 500 for internal server error.

Note: This code uses a global variable currentUser to keep track of the currently logged in user. It is important to ensure that this variable is thread-safe if multiple requests are handled concurrently.





Join Zoom Meeting
https://ufl.zoom.us/j/93037132618


