
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
## Front end work completed details
- Added modified journal entry page. This made us of an angular material stepper, with the first step being a stylized textarea where users can enter their journal entry for the day. As of now, we have a minimum word count of 500 for the entry section. A wordcount function updates and displays the user's wordcount as they type, and the stepper and the next button for the first step are disabled when the word count is less than 500. For demonstration purposes we have the minimum wordcount set to 50. When the user successfully submits their entry after the first step, a time-sensitive snackbar alerting them to choose a mood in the next step is displayed. The second step of the stepper is the "mood selection" step, where the user chooses between 5 "moods" that describe their day. The third step acts as reassurance for the user to submit their data, and on its button's click, the text entry and mood is sent to the backend.
-Added journal history page, which makes use of a calendar element to retrieve the logged-in user's entry from a date they select on the datepicker. The user can navigate dates using the calendar and display the text entry below for that day if there is one in the database. This also displays the chosen mood for that day.
-Moving forward, we hope to use the "mood" data to display the user's mood over time, possibly in graphs or charts. We also will highlight the dates for which a particular user has entries on the calendar to improve the user experience. 

-Created the general theme & design for the homepage. Gave the website a more branded feel.

-Simplified the concept for the user on the homepage. 

-Explained approaches users can take when writing journal entries.




## Frontend unit tests
-- Entry page tests
-Test to check the word count function on entry page
 it('should update the word count correctly', () => {
    component.text = 'Hello World, how are you?';
    component.updateWordCount();
    expect(component.wordCount).equal(5);
  });
- Next button should be disabled because word count is < 50
 it('should disable the matStepperNext button in the first step when the text is "Hello World, how are you?"', () => {
    const textarea = fixture.nativeElement.querySelector('textarea');
    textarea.value = 'Hello World, how are you?';
    textarea.dispatchEvent(new Event('input'));
    fixture.detectChanges();
    component.updateWordCount();
    fixture.detectChanges();
    const nextButton = fixture.nativeElement.querySelector('#nextButton');
    expect(nextButton.disabled).equal(true);
  });
- Next button should be enabled because word count is > 50
it('should enable the matStepperNext button in the first step when the text is greater than 50 words', () => {
    const textarea = fixture.nativeElement.querySelector('textarea');
    textarea.value = `Florida, our Alma Mater,thy glorious name we praise.
    All thy loyal sons and daughters,
    a joyous song shall raise.
    Where palm and pine are blowing,
    where southern seas are flowing,
    Shine forth thy noble gothic walls,
    thy lovely vine clad halls.
    Neath the orange and blue victorious, (throw right arm in the air)
    our love shall never fail.
    There's no other name so glorious, (throw right arm in the air)
    all hail, Florida, hail!`
    textarea.dispatchEvent(new Event('input'));
    fixture.detectChanges();
    component.updateWordCount();
    fixture.detectChanges();
    const nextButton = fixture.nativeElement.querySelector('#nextButton');
    expect(nextButton.disabled).equal(false);
  });
-- History page tests
- Checks that calendar defaults to current month when opened
it('calendar should default to MAR 2023 when page is opened', () => {
      cy.get('mat-calendar').click();
      cy.get('.mat-calendar-period-button').should('contain', 'MAR 2023');
});
- Checks that submit button is enabled/clickable on start
  it('submit button should be loaded and enabled on start', () => {
        cy.get('button').contains('Submit selection').click();
 });

## Backend unit tests
func TestJournalHandler(t *testing.T) : logs into existing account. Writes an entry. Stores entry under that user. 

func TestEntryRetrieverHandler() : Logs into existing account. Writes an entry. Stores entry under that user. Provides date for entry retrieval. Retrieves entry from that date.

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


