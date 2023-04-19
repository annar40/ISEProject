import { Component, NgModule, OnInit } from '@angular/core';
import { MoodChartComponent } from '../mood-chart/mood-chart.component';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-stats-page',
  templateUrl: './stats-page.component.html',
  styleUrls: ['./stats-page.component.css'],

})


export class StatsPageComponent{
  currentStreak: any;
  constructor(
    private router: Router,
    private httpClient: HttpClient,
    private activatedRoute: ActivatedRoute
  ) {}


  ngOnInit(): void {
    this.httpClient.get<any>('http://localhost:8000/retrieveDates').subscribe(data =>{
      console.log('Get streak: ', data.CurrentStreak);
      this.currentStreak = data.CurrentStreak;
      this.currentStreak = this.currentStreak +1;
    
    }, error  =>{
      console.log('Error getting streak', error);
    });

  }


}

