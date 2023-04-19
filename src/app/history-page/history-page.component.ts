import { Component } from '@angular/core';
import {MatDatepickerModule} from '@angular/material/datepicker';
import { MatDatepickerInputEvent } from '@angular/material/datepicker';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { formatDate } from '@angular/common';
import { MatCalendarCellCssClasses } from '@angular/material/datepicker';
import { ViewEncapsulation} from '@angular/core';

import {MatCalendarCellClassFunction} from '@angular/material/datepicker';

@Component({
  selector: 'app-history-page',
  templateUrl: './history-page.component.html',
  
  styleUrls: ['./history-page.component.css'],
  encapsulation: ViewEncapsulation.None,
})
export class HistoryPageComponent {
  selected!: Date | null;
  today: Date = new Date();
  journalEntry: any; // variable to store the journal entry response
  submittedDates: any;
  date: any;
  isDateLoaded: boolean = false;
  isHighLoaded: boolean = false;

  
  constructor(
    private router: Router,
    private httpClient: HttpClient,
    private activatedRoute: ActivatedRoute
  ) {}

  ngOnInit(): void {
    this.httpClient.get<any>('http://localhost:8000/retrieveDates').subscribe(data =>{
      console.log('Get dates: ', data.dates);
      this.submittedDates = data.dates;
      console.log('submitted dates', this.submittedDates);
      this.isDateLoaded = true;
      /*The boolean is set to true once the dates are returned from the backend so that the calendar can load first.
      The ngIf in the html file keeps the calendar from loading until it's true.
      */
    }, error  =>{
      console.log('Error getting dates', error);
    }
    );
  }


  
  dateClass: MatCalendarCellClassFunction<Date> = (cellDate, view) => {
    // Only highlight dates inside the month view.
    if (view === 'month') {
      const date = cellDate.getDate();
  
      // Highlight the dates in submittedDates.
      if (this.submittedDates && this.submittedDates.some((d: any) => d.date === formatDate(cellDate, 'yyyy-MM-dd', 'en'))) {
        return 'example-custom-date-class';
      }
    }
  
    return '';
  };
  
  
  
  
  
  

  submitDate() {
    if (this.selected) {
      const selectedDate = {
        date: this.selected.toLocaleDateString('en-CA', { year: 'numeric', month: '2-digit', day: '2-digit' })
      };
      
      const dateJson = JSON.stringify(selectedDate);
      console.log(dateJson);

      this.httpClient
        .post('http://localhost:8000/retrieveEntry', dateJson)
        .subscribe(
          (response) => {
            console.log('response', response);
            this.journalEntry = response; // store the response in the variable
          },
          (error: HttpErrorResponse) => {
            console.log('HTTP error status:', error.status);
            // only redirect if the error status is not 200 OK
            if (error.status === 200) {
              console.log('Journal Entry Retrieved');
              this.router.navigate(['../', 'entry'], {
                relativeTo: this.activatedRoute,
              });
            }
          }
        );
    } else {
      //a date isn't selected
    }
  }

}
