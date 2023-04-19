import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { StatsPageComponent } from './stats-page.component';
import 'chart.js/auto';
import { Chart } from 'chart.js';


//Chart.register(...registerables);

describe('Stats Page', () => {
  let component: StatsPageComponent;
  let fixture: ComponentFixture<StatsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [StatsPageComponent],
      imports: [HttpClientTestingModule],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(StatsPageComponent);
    component = fixture.componentInstance;

    component.currentStreak = 60;
    component.totalEntries = 120;
    component.testMoods = [
        { name: "😁 Great", value: 60 },
        { name: "🙂 Good", value: 50 },
        { name: "🙃 Ok", value: 40 },
        { name: "😔 Sad", value: 30 },
        { name: "😡 Angry", value: 20 },
        { name: "😰 Anxious", value: 10 }
      ];

      const xValues = component.testMoods.map(mood => mood.name);
      const yValues = component.testMoods.map(mood => mood.value);
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
    fixture.detectChanges();
  });

  it('displays the current streak value', () => {
    cy.get('.large-text').contains('60').should('be.visible');
  });
  it('displays the total entries value', () => {
    cy.get('.large-text').contains('120').should('be.visible');
  });
  it('displays the pie chart', () => {
    cy.get('#myChart').should('be.visible');
  });



});
