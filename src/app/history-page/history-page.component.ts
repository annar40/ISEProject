import { Component } from '@angular/core';
import {MatDatepickerModule} from '@angular/material/datepicker';
import { MatDatepickerInputEvent } from '@angular/material/datepicker';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';


@Component({
  selector: 'app-history-page',
  templateUrl: './history-page.component.html',
  styleUrls: ['./history-page.component.css']
})
export class HistoryPageComponent {
  selected!: Date | null;

  constructor(
    private router: Router,
    private httpClient: HttpClient,
    private activatedRoute: ActivatedRoute
  ) {}

  submitDate(){
    if(this.selected){
      //a date is selected do something, submit logic here
      console.log(this.selected);

      const selectedDate = { date: this.selected.toISOString().slice(0, 10) };

      this.httpClient
        .post('http://localhost:8000/retrieveEntry', selectedDate)
        .subscribe(
          (response) => {
            console.log('response', response);
          },
          (error: HttpErrorResponse) => {
            console.log('HTTP error status:', error.status);
            // only redirect if the error status is not 200 OK
            if (error.status === 200) {
              console.log('Journal Entry Retrieved');
              this.router.navigate(['../', 'journal'], {
                relativeTo: this.activatedRoute,
              });
            }
          }
        );
    }
    else{
      //a date isn't selected
    }
  }

}
