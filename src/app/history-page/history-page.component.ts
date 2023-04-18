import { Component } from '@angular/core';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatDatepickerInputEvent } from '@angular/material/datepicker';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { formatDate } from '@angular/common';
import { MatCalendarCellCssClasses } from '@angular/material/datepicker';

@Component({
  selector: 'app-history-page',
  templateUrl: './history-page.component.html',
  styleUrls: ['./history-page.component.css']
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
  ) { }

  ngOnInit(): void {
    this.getDates();
  }

  getDates(): void {
    this.httpClient.get<any>('http://localhost:8000/retrieveDates').subscribe(data => {
      const dates = data.dates;
      console.log('Get dates: ', dates);
      this.submittedDates = dates.flat().map((dateObj: { date: string }) => dateObj.date);
      // console.log('Submitted dates', this.submittedDates);

      console.log('Date format', new Date(this.submittedDates[0]).toLocaleDateString());

      this.isDateLoaded = true;
      console.log('Is date loaded?', this.isDateLoaded);

      // Highlight all the dates
      this.submittedDates.forEach((dateStr: string) => {
        const workPls = new Date(dateStr);
        this.highlightDate(workPls);
      });

    }, error => {
      console.log('Error getting dates', error);

    });
  }

  highlightDate(date: Date | null): string {
    console.log('is date loaded?', this.isDateLoaded);

    if (date === null || !this.isDateLoaded) {
      return '';
    }
    console.log('is date loaded?', this.isDateLoaded);
  
    const formattedDate = date.toLocaleDateString('en-CA', { year: 'numeric', month: '2-digit', day: '2-digit'});
    console.log('formatted', formattedDate);
    console.log(this.submittedDates);

    console.log('dateLoaded', this.isDateLoaded);
    const isHighlighted = this.submittedDates.includes(formattedDate);

    console.log("highlighted ", isHighlighted)

    return isHighlighted ? 'highlighted-date' : '';
  }


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
