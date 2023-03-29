import { ComponentFixture, TestBed } from '@angular/core/testing';
import { EntryPageComponent } from './entry-page.component';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { HttpClientModule } from '@angular/common/http';
import { RouterTestingModule } from '@angular/router/testing';
import { MatStepperModule } from '@angular/material/stepper';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';


describe('EntryPageComponent', () => {
  let component: EntryPageComponent;
  let fixture: ComponentFixture<EntryPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [EntryPageComponent],
      imports: [MatSnackBarModule, HttpClientModule, RouterTestingModule, MatStepperModule,FormsModule,BrowserAnimationsModule],
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

  it('should disable the matStepperNext button in the first step when the text is "Hello World, how are you?"', () => {
    const textarea = fixture.nativeElement.querySelector('textarea');
    textarea.value = 'Hello World, how are you?';
    textarea.dispatchEvent(new Event('input'));
    fixture.detectChanges();
    component.updateWordCount();
    fixture.detectChanges();
    const nextButton = fixture.nativeElement.querySelector('#nextButton');
    expect(nextButton.disabled).equal(true);
  });

  it('should enable the matStepperNext button in the first step when the text is "Hello World, how are you?"', () => {
    const textarea = fixture.nativeElement.querySelector('textarea');
    textarea.value = `Florida, our Alma Mater,thy glorious name we praise.
    All thy loyal sons and daughters,
    a joyous song shall raise.
    Where palm and pine are blowing,
    where southern seas are flowing,
    Shine forth thy noble gothic walls,
    thy lovely vine clad halls.
    Neath the orange and blue victorious, (throw right arm in the air)
    our love shall never fail.
    There's no other name so glorious, (throw right arm in the air)
    all hail, Florida, hail!`
    textarea.dispatchEvent(new Event('input'));
    fixture.detectChanges();
    component.updateWordCount();
    fixture.detectChanges();
    const nextButton = fixture.nativeElement.querySelector('#nextButton');
    expect(nextButton.disabled).equal(false);
  });
  

});
