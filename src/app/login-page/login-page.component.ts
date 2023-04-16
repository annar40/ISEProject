import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, Validators, FormControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Injectable } from '@angular/core';

@Injectable()
export class MyService {
  constructor(private http: HttpClient) {}

  retrieveDates() {
    return this.http.get('/retrieveDates');
  }
}

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent{
  loginForm: FormGroup = new FormGroup(
    {
      name: new FormControl('', [Validators.required]),
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [
        Validators.required,
        Validators.minLength(5),
      ]),
    },
    
  );

  constructor(
    private router: Router,
    private httpClient: HttpClient,
    private activatedRoute: ActivatedRoute
  ) {}

   ngOnInit(): void {}

  onSubmit() {
    console.log( this.loginForm.value);

    this.httpClient.post(
        'http://localhost:8000/login',
        
        this.loginForm.value
      )
      .subscribe(
        (response) => {
          console.log('response', response);

          // Call the new method here
          
          
        },
        (error: HttpErrorResponse) => {
          console.log('HTTP error status:', error.status);
          // only redirect if the error status is not 200 OK
          if (error.status === 200) {
            this.loginForm.reset();
            this.sendNewRequest();
            console.log('Signup Successful');
            this.router.navigate(['../', 'entry'], {
              relativeTo: this.activatedRoute,

            });
          }
        }
      );
      
  }

  // New method to send another POST request
  sendNewRequest() {
    this.httpClient.post(
        'http://localhost:8000/retrieveDates',
        
        // Pass any data you want to send in the request body
        {someData: 'someValue'}
      )
      .subscribe(
        (response) => {
          console.log('New response', response);
          
        },
        (error: HttpErrorResponse) => {
          console.log('New HTTP error status:', error.status);
        }
      );
  }
}
2