import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators , FormControl} from '@angular/forms';
import { DatePipe } from '@angular/common';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ViewChild } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { MatStep, MatStepper } from '@angular/material/stepper';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-entry-page',
  templateUrl: './entry-page.component.html',
  styleUrls: ['./entry-page.component.css']
})
export class EntryPageComponent {
  @ViewChild('firstStep') firstStep!: MatStep;
  @ViewChild('Stepper') Stepper!: MatStepper;

  n = new Date();
  constructor(public snackBar: MatSnackBar,
    private router: Router,
    private httpClient: HttpClient,
    private activatedRoute: ActivatedRoute) { }
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


  
  journalEntry: string = '';


  ngOnInit(): void {}

  onSubmit() {
    const requestBody = { text: this.text };
    console.log(requestBody);

    this.httpClient.post('http://localhost:8000/journalEntry', JSON.stringify(requestBody))
      .subscribe(
        (response) => {
          console.log('response', response);
          
        },
        (error: HttpErrorResponse) => {
          console.log('HTTP error status:', error.status);
          // only redirect if the error status is not 200 OK
          if (error.status === 200) {
            
            console.log('Journal Entry Stored');
            this.router.navigate(['../', 'history'], {
              relativeTo: this.activatedRoute,
            });
          }
        }
      );
  }
  
}
