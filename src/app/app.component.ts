import { Component } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SignupPageComponent } from './signup-page/signup-page.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { HttpClient } from '@angular/common/http';
import { NgModule } from '@angular/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'morning-pages-0';
  username = '';
  password = '';
  message = '';

  email = '';

  login() {
    console.log('Login button clicked');
  }

  constructor(private http: HttpClient) {}

  submitForm() {
    const url = 'http://localhost:8080/signup'; // change this to your server's URL
    const body = { username: this.username, password: this.password };

    this.http.post(url, body).subscribe(
      (response: any) => {
        this.message = 'Signup successful!';
      },
      (error: any) => {
        console.log(error);
        this.message = 'Signup failed. Please try again.';
      }
    );
  }
}

const routes: Routes = [
  { path: 'signup', component: SignupPageComponent },
  { path: 'login', component: LoginPageComponent },
  { path: '', redirectTo: '/login', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
