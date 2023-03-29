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
        fixture.detectChanges();
      });
    it('calendar should default to MAR 2023 when page is opened', () => {
      cy.get('mat-calendar').click();
      cy.get('.mat-calendar-period-button').should('contain', 'MAR 2023');
    });

    it('submit button should be loaded and enabled on start', () => {
        cy.get('button').contains('Submit selection').click();
    });

  });
  