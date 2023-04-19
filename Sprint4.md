# WORK TO COMPLETE:

Make progress on issues uncompleted in Sprint 3, or new issues discovered during Sprint 3.

Write test for new functionality implemented. 

A front-page readme that details requirements for running and using your application.

Finish working on the homepage that was started in Sprint 3

## things to still do
- update the readme on running details
- create ui message for failed log in attempt?
- frontend streaks

# SUBMISSIONS:
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
- Created a handler that retrieves all of the dates in which the user has a stored entry. This handler is used to send the frontend the user's streak.
- Updated the journal handler to adjust the user's streak upon submitting an entry. 
### Frontend:
- Updated calendar widget to highlight the dates of stored entries.
- Added stats page that displays the user's streak.
# Tests
### Backend:
- func TestForDuplicateUsers(t * testing.T) : attempts to create a new account with a username that exists in database. 
- func TestRetrieveDatesHandler(t * testing.T) : logs into existing user, calls retrieve date handler, checks for the entry dates.
- Modified func TestRetrieveEntryHandler(t * testing.T) : selects a date that does not contain a stored entry. Expects a null response.


### Frontend:
List frontend unit and Cypress tests


# Updated API

