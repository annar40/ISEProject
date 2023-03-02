
import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
//import{HttpClient} from '@angular/common/http/testing';
import { HttpClientModule } from '@angular/common/http';
import {Router } from '@angular/router';
import { RouterTestingModule } from "@angular/router/testing";
import { LoginPageComponent } from './login-page.component';
import { BrowserModule, By } from '@angular/platform-browser';
import { MatCardModule } from '@angular/material/card';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatFormFieldModule } from '@angular/material/form-field';

describe('LoginPageComponent', () => {
  let comp: LoginPageComponent;
  let fixture: ComponentFixture<LoginPageComponent>;


  beforeEach(async () => {
     await TestBed.configureTestingModule({
      imports:[FormsModule, BrowserModule, HttpClientModule, ReactiveFormsModule, RouterTestingModule, MatCardModule
      ,MatToolbarModule,MatFormFieldModule,],
      declarations: [ LoginPageComponent ],
    }).compileComponents()
  });



  it('form should be valid -- all login fields filled', async() => {
  
    fixture = TestBed.createComponent(LoginPageComponent);
    comp = fixture.componentInstance;

    comp.loginForm.controls['name'].setValue('John');
    comp.loginForm.controls['email'].setValue('john@gmail.com');
    comp.loginForm.controls['password'].setValue('abc123');
    expect(comp.loginForm.valid).toBeTruthy();
  });

  it('form should be invalid -- name field is empty', async() => {
  
    fixture = TestBed.createComponent(LoginPageComponent);
    comp = fixture.componentInstance;

    comp.loginForm.controls['name'].setValue('');
    comp.loginForm.controls['email'].setValue('john@gmail.com');
    comp.loginForm.controls['password'].setValue('abc123');
    expect(comp.loginForm.valid).toBeFalsy();
  });

  
  it('form should be invalid -- password field is too short', async() => {
  
    fixture = TestBed.createComponent(LoginPageComponent);
    comp = fixture.componentInstance;

    comp.loginForm.controls['name'].setValue('John');
    comp.loginForm.controls['email'].setValue('john@gmail.com');
    comp.loginForm.controls['password'].setValue('abc');
    expect(comp.loginForm.valid).toBeFalsy();
  });

});

