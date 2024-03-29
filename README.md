# MorningPages

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 15.1.4.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

## Running the Application Step by Step

1.) Make sure to install node.js and Golang

2.) Run `npm install` to install all libraries and dependencies needed for the code. 

3.) In the command line navigate to the folder of the project and run `ng serve` to compile the angular code 

4.) In a new command line, navigate to the same directory and type `go run hello.go` to compile the backend API .  Select "Allow Access" on the Window Security Alert pop-up.

5.) Navigate to `http://localhost:4200/signup` to get started with our application

## Using the application
1.) Create an account on the signup page or login to an existing account

2.) Once logged in, you'll be redirected to add an entry on the entry page. There is a 500 word minimum for the text entry. In the second step, you can choose a mood that will be logged along with the date and the text entry.

3.) Navigate to the history page, where you can view your past submissions via the calendar. Dates which have an entry are higlighted in purple on the calendar. The retrieved entry will show the date, mood, and text for the day you select.

4.) On the stats page, you can see your current streak (the number of consecutive days for which you have an entry), your total entries, and a pie chart which shows the number of times you submitted an entry with each mood.

## Running the frontend tests
On the finalFrontTesting branch, with cypress installed run `cypress open`

## Running the backend tests
In the command line navigate to the folder of the project and run `go test` 

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.


## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.


