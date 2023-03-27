import { Component } from '@angular/core';
import {MatDatepickerModule} from '@angular/material/datepicker';
import { MatDatepickerInputEvent } from '@angular/material/datepicker';
@Component({
  selector: 'app-history-page',
  templateUrl: './history-page.component.html',
  styleUrls: ['./history-page.component.css']
})
export class HistoryPageComponent {
  selected!: Date | null;
  submitDate(){
    if(this.selected){
      //a date is selected do something, submit logic here
      console.log(this.selected);
    }
    else{
      //a date isn't selected
    }
  }
}
