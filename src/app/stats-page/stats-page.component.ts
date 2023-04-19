import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Chart, registerables } from 'chart.js';
import { ActivatedRoute,Router } from '@angular/router';

Chart.register(...registerables);

@Component({
  selector: 'app-stats-page',
  templateUrl: './stats-page.component.html',
  styleUrls: ['./stats-page.component.css']
})
export class StatsPageComponent implements OnInit {
  @ViewChild('chartCanvas') chartCanvas!: ElementRef;
  currentStreak: any;
  moods: any;
  totalEntries: any;


  constructor(
    private router: Router,
    private httpClient: HttpClient,
    private activatedRoute: ActivatedRoute
  ) {}

  createEntry() {
    this.router.navigate(['/entry']);
  }

  ngOnInit(): void {
    this.httpClient.get<any>('http://localhost:8000/retrieveDates').subscribe(data => {
      console.log('Get dates', data.dates);
      this.totalEntries = data.dates.length;
      console.log('Total entries', this.totalEntries);
      console.log('Get streak:', data.CurrentStreak);
      this.currentStreak = data.CurrentStreak;
      this.currentStreak = this.currentStreak ;
    }, error => {
      console.log('Error getting streak:', error);
    });

    this.httpClient.get<any>('http://localhost:8000/retrieveMoods').subscribe(data => {
      console.log('Get mood:', data.moods);
      const xValues = Object.keys(data.moods);
      const yValues = Object.values(data.moods);
      const barColors = [
        '#EAEAEA',
        '#CBC5EA',
        '#73628A',
        '#1F2C4A',
        '#190320',
        '#08060B'
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
        options: {}
      });
    }, error => {
      console.log('Error getting moods:', error);
    });
  }
}
