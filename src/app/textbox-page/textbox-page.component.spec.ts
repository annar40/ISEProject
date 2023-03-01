import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TextboxPageComponent } from './textbox-page.component';

describe('TextboxPageComponent', () => {
  let component: TextboxPageComponent;
  let fixture: ComponentFixture<TextboxPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TextboxPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TextboxPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

