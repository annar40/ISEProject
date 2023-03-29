import { ComponentFixture, TestBed } from '@angular/core/testing';
import { EntryPageComponent } from './entry-page.component';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { HttpClientModule } from '@angular/common/http';
import { RouterTestingModule } from '@angular/router/testing';

describe('EntryPageComponent', () => {
  let component: EntryPageComponent;
  let fixture: ComponentFixture<EntryPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [EntryPageComponent],
      imports: [MatSnackBarModule, HttpClientModule, RouterTestingModule],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EntryPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should update the word count correctly', () => {
    component.text = 'Hello World, how are you?';
    component.updateWordCount();
    expect(component.wordCount).equal(5);
  });
});
