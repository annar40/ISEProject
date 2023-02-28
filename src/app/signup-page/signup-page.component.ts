import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, Validators, FormControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';




@Component({
  selector: 'app-signup-page',
  templateUrl: './signup-page.component.html',
  styleUrls: ['./signup-page.component.css']
})
export class SignupPageComponent implements OnInit {

  

  signupForm: FormGroup = new FormGroup(
    {
      name: new FormControl('', [Validators.required]),
      lastName: new FormControl('', [Validators.required]),
      phoneNumber: new FormControl('', [Validators.required]),
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [
        Validators.required,
        Validators.minLength(5),
      ]),
      confirmPassword: new FormControl('', [Validators.required]),
    },
    
  );

  constructor(
    private router: Router,
    private httpClient: HttpClient,
    private activatedRoute: ActivatedRoute
  ) {}

   ngOnInit(): void {}

  onSubmit() {
    console.log( this.signupForm.value);

    this.httpClient
      .post(
        'https://thoughtdump-4b31d-default-rtdb.firebaseio.com/users.json',
        this.signupForm.value
      )
      .subscribe(
        (response) => {
          console.log('response', response);
          this.signupForm.reset();
          this.router.navigate(['../', 'login'], {
            relativeTo: this.activatedRoute,
          });
        },
        (error) => {
          console.log(error);
        }
      );
  }
}