import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent {
  username = '';
  password = '';
  message = '';

  email = '';

  login()
  {
    console.log('Login button clicked');
  }
  constructor(private http: HttpClient) {}

  // submitForm() {
  //   const url = 'http://localhost:8080/signup'; // change this to your server's URL
  //   const body = { username: this.username, password: this.password };

  //   this.http.post(url, body).subscribe(
  //     (response: any) => {
  //       this.message = 'Signup successful!';
  //     },
  //     (error: any) => {
  //       console.log(error);
  //       this.message = 'Signup failed. Please try again.';
  //     }
  //   );
  // }
}
