import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DatePipe } from '@angular/common';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ViewChild } from '@angular/core';
import { MatStep, MatStepper } from '@angular/material/stepper';

@Component({
  selector: 'app-entry-page',
  templateUrl: './entry-page.component.html',
  styleUrls: ['./entry-page.component.css']
})
export class EntryPageComponent {
  @ViewChild('firstStep') firstStep!: MatStep;
  @ViewChild('Stepper') Stepper!: MatStepper;

  n = new Date();
  constructor(public snackBar: MatSnackBar) { }
  text: string = '';
  wordCount: number = 0;

  updateWordCount() {
    this.wordCount = this.text.trim().split(' ').length;
  }
  onEndFirst(){
    this.firstStep.completed = true;
    this.Stepper.next();
    this.snackBar.open('Your journal entry was logged! Now choose a mood for the day', 'Ok', { duration: 3000 })
  }

  
}
