import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, Validators, FormControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';


@Component({
  selector: 'app-textbox-page',
  templateUrl: './textbox-page.component.html',
  styleUrls: ['./textbox-page.component.css']
})
export class TextboxPageComponent {

  journalForm: FormGroup = new FormGroup({
    journalEntry: new FormControl('', [Validators.required]),
  });

  constructor(
    private router: Router,
    private httpClient: HttpClient,
    private activatedRoute: ActivatedRoute
  ) {}

  ngOnInit(): void {}

  onSubmit() {
    console.log(this.journalForm.value);

    this.httpClient
      .post('http://localhost:8000/journalEntry', this.journalForm.value)
      .subscribe(
        (response) => {
          console.log('response', response);
          
        },
        (error: HttpErrorResponse) => {
          console.log('HTTP error status:', error.status);
          // only redirect if the error status is not 200 OK
          if (error.status === 200) {
            this.journalForm.reset();
            console.log('Journal Entry Stored');
            this.router.navigate(['../', 'journal'], {
              relativeTo: this.activatedRoute,
            });
          }
        }
      );
  }
}
const clearbutton = document.getElementById("clearbutton");