import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
//import{HttpClient} from '@angular/common/http/testing';
import { HttpClientModule } from '@angular/common/http';
import {Router } from '@angular/router';
import { RouterTestingModule } from "@angular/router/testing";
import { SignupPageComponent } from './signup-page.component';
import { BrowserModule, By } from '@angular/platform-browser';
import { MatCardModule } from '@angular/material/card';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatFormFieldModule } from '@angular/material/form-field';


describe('SignUpButton', ()=>{
    it('can mount', () =>{
        cy.mount(SignupPageComponent);
    })
});