import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-signup-page',
  templateUrl: './signup-page.component.html',
  styleUrls: ['./signup-page.component.css']
})
export class SignupPageComponent {
  firstName = '';
  lastName = '';
  email = '';
  password = '';
  confirmPassword = '';
  message = '';

  constructor(private http: HttpClient) {}

  createAccount() {
    const userData = {
      firstName: this.firstName,
      lastName: this.lastName,
      email: this.email,
      password: this.password,
      confirmPassword: this.confirmPassword
    };
   
    this.http.post('/sign-up', userData).subscribe((response) => {
      console.log(response);
    //           this.message = 'Signup successful!';
    //   },
    //   (error: any) => {
    //     console.log(error);
    //     this.message = 'Signup failed. Please try again.';
      }
    );
  }
  submitForm() {
    this.createAccount();
  }
}
