
## things to still do
- A front-page readme that details requirements for running and using your application.

# SUBMISSIONS (video info):
Submission Format: GitHub & Video Links (Use comments on submission page for multiple links)

Narrated video presentation. Split the presentation such that each member of your team narrates a portion. 

Presentation should include:
- Demonstrate new functionality implemented.
- Show results of all unit tests (including those from Sprint 3).
- Finally, give an overview of your completed project as if you were pitching it to someone who has never seen it:
- Demonstrate all front-end functionality
- Detailed explanation of backend API

# Work Accomplished
### Bug Fixes: 
- Resolved Calendar bug: User can now submit a date that does not have an entry stored without any backend crashes.
- Resolved Duplicate User bug: User can no longer create an account with a username that exists the data base. The email and password are no longer overwritten. User is prompted with error message "User already exists."
### Backend:
- Modified user struct to include a streak characteristic.
- Created a handler that retrieves all of the dates in which the user has a stored entry. This handler is used to send the frontend the user's streak.
- Updated the journal handler to adjust the user's streak upon submitting an entry. 
- Created a handler that retrieves all of the moods from all of the user's stored entries. This handler is used to send the frontend data for stats pie chart.
### Frontend:
- Updated calendar widget to highlight the dates of stored entries.
- Created stats page that displays the user's streak, total entries, and pie chart of all previous moods.
- Implemented routing for all buttons on homepage.
- Added history tab to navigation bar.
# Tests
### Backend:
- func TestForDuplicateUsers(t * testing.T) : attempts to create a new account with a username that exists in database. 
- func TestRetrieveDatesHandler(t * testing.T) : logs into existing user, calls retrieve date handler, checks for the entry dates.
- Modified func TestRetrieveEntryHandler(t * testing.T) : selects a date that does not contain a stored entry. Expects a null response.
- func TestRetrieveMoodsHandler(t * testing.T) : logs into existing user, calls retrieve mood handler, checks for the entry moods.


### Frontend:
-     it('should disable tomorrow (4/20) on the calendar', () => {
        cy.get('mat-calendar').click();
        cy.get('.mat-calendar-body-cell-content').contains('20').parent().should('have.class', 'mat-calendar-body-disabled');
    });
    test to check if tomorrow's date is disabled (all future dates on the calendar of history page should be disabled)
-  it('displays the current streak value', () => {
    cy.get('.large-text').contains('60').should('be.visible');
  });
  checks that the streak value when its variable is set to 60 is displaying properly
-  it('displays the total entries value', () => {
    cy.get('.large-text').contains('120').should('be.visible');
  });
  checks that the total entries value when its variable is set to 120 is displaying properly
-  it('displays the pie chart', () => {
    cy.get('#myChart').should('be.visible');
  });
  checks that the pie chart is displayed when created


# Updated API
The API provides several endpoints for user authentication and journal entries. The endpoints are listed below:

1.     POST /signup - Create a new user account
This endpoint allows a user to sign up. The request body should contain a JSON object with the following fields: { "name": "string", "email": "string", "password": "string" } The name field is used as the document ID in Firestore.

Response:
- Status code 200 on success
- Status code 400 on invalid form data
- Status code 409 if username is already taken
- Status code 500 on any other error

&nbsp;

2.     POST /login - Login a user
This endpoint allows a user to login. The request body should contain a JSON object with the following fields: { "name": "string", "email": "string", "password": "string" }

The name field is used as the document ID in Firestore

Response:

- Status code 200 on success
- Status code 400 on invalid form data
- Status code 401 if email or password is incorrect
- Status code 500 on any other error

&nbsp;

3.     POST /journalEntry - Add a journal entry
This endpoint allows a user to create a journal entry. The request body should contain a JSON object with the following fields: { "text": "string", "mood": "string" } The text field is the journal entry text and the mood field is the mood for that entry.

Response:

- Status code 200 on success
- Status code 400 on invalid form data
- Status code 500 on any other error

&nbsp;

4.     POST /retrieveEntry - Retrieve a journal entry
This endpoint allows a user to retrieve a journal entry for a specific date. The request body should contain a JSON object with the following field: { "date": "string" } The date field should be in the format "yyyy-mm-dd".

Response:

- Status code 200 on success
- Status code 400 on invalid form data
- Status code 500 on any other error

&nbsp;

5.     POST /retrieveDates - Retrieve all dates with journal entries
This endpoint allows a user to retrieve all dates with journal entries for the logged in user. 

Response:
- Status code 200 on success
- Status code 500 on any other error

&nbsp;

6.     POST /retrieveMoods - Retrieve all moods from all stored journal entries
This endpoint allows a user to retrieve all moods ever stored of the logged in user. 

Response:
- Status code 200 on success
- Status code 500 on any other error

All endpoints support Cross-Origin Resource Sharing (CORS).
