import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HistoryPageComponent } from './history-page.component';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { HttpClientModule, HttpInterceptor } from '@angular/common/http';
import { RouterTestingModule } from '@angular/router/testing';
import { MatStepperModule } from '@angular/material/stepper';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';

describe('History Page', () => {
    let component:HistoryPageComponent;
    let fixture:ComponentFixture<HistoryPageComponent>;
    beforeEach(async () => {
        await TestBed.configureTestingModule({
          declarations: [HistoryPageComponent],
          imports: [MatSnackBarModule, HttpClientModule, RouterTestingModule, MatStepperModule,FormsModule,BrowserAnimationsModule, MatDatepickerModule, MatNativeDateModule,],
        }).compileComponents();
      });
    
      beforeEach(() => {
        fixture = TestBed.createComponent(HistoryPageComponent);
        component = fixture.componentInstance;
        component.submittedDates = ['2023-04-19', '2023-04-13', '2022-04-04'];

        component.isDateLoaded = true;
          
        fixture.detectChanges();
      });
    it('calendar should default to APR 2023 when page is opened', () => {
      cy.get('mat-calendar').click();
      cy.get('.mat-calendar-period-button').should('contain', 'APR 2023');
    });

    it('submit button should be loaded and enabled on start', () => {
        cy.get('button').contains('Submit selection').click();
    });
    it('should disable tomorrow (4/20) on the calendar', () => {
        cy.get('mat-calendar').click();
        cy.get('.mat-calendar-body-cell-content').contains('20').parent().should('have.class', 'mat-calendar-body-disabled');
    });
      
      
      


  });
  