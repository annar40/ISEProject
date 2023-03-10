


## Video
https://youtu.be/T0HGBoPXOaA


## Work Completed in Sprint 2

We had many troubles connecting our frontend and backend together. 

The golang file was completely rewritten multiple times until a successful integration occured.
Forms were added to the sign up and login typescrpit components to send the information to the golang server which parsed the json and sent it into the firestore database. 

Once the golang and angular were integrated, backend unit tests were written for the login and signup functions alongside front end test.

A navigation bar was added along with buttons for routing: home, log in, sign up, and about

The journaling page was created, along with components for filling out the about and home pages in the future. The sign up page routes users to the login page upon account creation. The login page routes users to the journal page upon successful login.  

If the login information is not completely filled out or is invalid, the "sign in" button will appear grey and unuasble.



## Unit Tests for Frontend
1. Password must be at least 5 characters.

  it('form should be invalid -- password field is too short', async() => {
  
    fixture = TestBed.createComponent(LoginPageComponent);
    comp = fixture.componentInstance;

    comp.loginForm.controls['name'].setValue('John');
    comp.loginForm.controls['email'].setValue('john@gmail.com');
    comp.loginForm.controls['password'].setValue('abc');
    expect(comp.loginForm.valid).toBeFalsy();
  });

2. Name field is empty
it('form should be invalid -- name field is empty', async() => {
  
    fixture = TestBed.createComponent(LoginPageComponent);
    comp = fixture.componentInstance;

    comp.loginForm.controls['name'].setValue('');
    comp.loginForm.controls['email'].setValue('john@gmail.com');
    comp.loginForm.controls['password'].setValue('abc123');
    expect(comp.loginForm.valid).toBeFalsy();
  });

3.  All components must be filled for login to true

  it('form should be valid -- all login fields filled', async() => {
  
    fixture = TestBed.createComponent(LoginPageComponent);
    comp = fixture.componentInstance;

    comp.loginForm.controls['name'].setValue('John');
    comp.loginForm.controls['email'].setValue('john@gmail.com');
    comp.loginForm.controls['password'].setValue('abc123');
    expect(comp.loginForm.valid).toBeTruthy();
  });


## Cypress Test
1. Check for mounting success of navigation bar

describe('MainNavbarComponent', ()=>{
    it('can mount', () =>{
        cy.mount(MainNavbarComponent);
    })
});


2. Check for a valid click on first button (home button) on the navigation bar
 
describe('MainNavBarComponent', () => {
    
    it('expected a logged successful click on home', () => {
      cy.mount(MainNavbarComponent);
      cy.get('a').eq(0).click();
    });
  });



## Unit Tests for Backend
func TestSignupHandler(t *testing.T) -checks for server status 200  and message "Login Successful"

func TestLoginHandler(t *testing.T) -checks for server status 200 and message "User data written to Firestore"


## **API Documentation:**

API Documentation: Package Name: Firebase Auth API Package Description: This package provides APIs to handle user authentication using Firebase and Firestore.

**Imported Packages:**

context: This package is used for passing context between functions. encoding/json: This package is used to encode and decode JSON data. fmt: This package is used for formatting output. log: This package is used for logging error messages. net/http: This package is used to implement HTTP client and server. cloud.google.com/go/firestore: This package is used to interact with Firestore. firebase.google.com/go: This package is used to initialize Firebase app. github.com/rs/cors: This package is used to handle Cross-Origin Resource Sharing (CORS). google.golang.org/api/option: This package is used to provide options for authentication and other settings.

**Structures:**

User: This structure is used to represent user information. It contains four fields: Name: string, the name of the user. Phone: string, the phone number of the user. Email: string, the email address of the user. Password: string, the password of the user.

**Variables:**

ctx: context.Context, a context for passing context between functions.

**Functions:**

main(): This function is the entry point of the program. It initializes the Firebase app and creates a Firestore client. Then it attaches signup and login handlers to the HTTP server and starts the server.

signupHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request): This function is used to handle user signup requests. It takes a Firestore client as an input parameter and returns a function that handles HTTP requests. This function parses form data, writes user data to Firestore, and sends a success response.

loginHandler(client *firestore.Client) func(w http.ResponseWriter, r *http.Request): This function is used to handle user login requests. It takes a Firestore client as an input parameter and returns a function that handles HTTP requests. This function parses form data, retrieves user data from Firestore, checks if email and password match, and sends a success response if the login is successful.

**API Endpoints:**

/signup: This endpoint is used for user signup requests. It accepts a POST request with user data in JSON format. The user data must contain the name, email, and password fields. If the signup is successful, it sends a 200 OK response with the message "User data written to Firestore".

/login: This endpoint is used for user login requests. It accepts a POST request with user data in JSON format. The user data must contain the name, email, and password fields. If the login is successful, it sends a 200 OK response with the message "Login Successful".

Sample Usage: To use this package, you can import it in your Go program and call the main function. You can then send HTTP requests to the /signup and /login endpoints to handle user signup and login requests.
