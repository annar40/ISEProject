import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, Validators, FormControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-signup-page',
  templateUrl: './signup-page.component.html',
  styleUrls: ['./signup-page.component.css']
})
export class SignupPageComponent {
  usernameTakenError = false;
  signupForm: FormGroup = new FormGroup({
    name: new FormControl('', [Validators.required]),
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [
      Validators.required,
      Validators.minLength(5),
    ]),
    confirmPassword: new FormControl('', [Validators.required]),
  });

  constructor(
    private router: Router,
    private httpClient: HttpClient,
    private activatedRoute: ActivatedRoute
  ) {}

  ngOnInit(): void {}

  onSubmit() {
    console.log(this.signupForm.value);

    this.httpClient
      .post('http://localhost:8000/signup', this.signupForm.value)
      .subscribe(
        (response) => {
          console.log('response', response);
        },
        (error: HttpErrorResponse) => {
          console.log('HTTP error status:', error.status);
          // only redirect if the error status is not 200 OK
          if (error.status === 200) {
            this.signupForm.reset();
            console.log('Signup Successful');
            this.router.navigate(['../', 'login'], {
              relativeTo: this.activatedRoute,
            });
          }
          else if(error.status === 409)
          {
            //user name already taken
            this.usernameTakenError = true;
          }
        }
      );
  }
}
