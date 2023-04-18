import { Component, OnInit } from '@angular/core';
import { Chart, registerables } from 'chart.js';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
Chart.register(...registerables);

@Component({
  selector: 'app-mood-chart',
  templateUrl: './mood-chart.component.html',
  styleUrls: ['./mood-chart.component.css']
})
export class MoodChartComponent implements OnInit{
  okCount: number = 0;
  moodString: string = "";

  constructor(
    private router: Router,
    private httpClient: HttpClient,
    private activatedRoute: ActivatedRoute
  ) {}
  ngOnInit(): void {
    this.httpClient.get<any>('http://localhost:8000/retrieveDates').subscribe(data =>{
      console.log('Get entry: ', data.dates);
      //this.moodString = data.mood;
      //console.log('submitted dates', this.submittedDates);
      //this.isDateLoaded = true;
      /*The boolean is set to true once the dates are returned from the backend so that the calendar can load first.
      The ngIf in the html file keeps the calendar from loading until it's true.
      */
    }, error  =>{
      console.log('Error getting mood', error);
    }
    );

    var xValues = ["Italy", "France", "Spain", "USA", "Argentina"];
    var yValues = [55, 49, 44, 24, 15];
    var barColors = [
      "#EAEAEA",
      "#CBC5EA",
      "#73628A",
      "#1F2C4A",
      "#190320"
    ];
    
    new Chart("myChart", {
      type: "pie",
      data: {
        labels: xValues,
        datasets: [{
          backgroundColor: barColors,
          data: yValues
        }]
      },
      options: {
      }
    });
  }
}
