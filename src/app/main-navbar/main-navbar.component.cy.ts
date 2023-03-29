import { MainNavbarComponent } from "./main-navbar.component";

describe('MainNavbarComponent', ()=>{
    it('can mount', () =>{
        cy.mount(MainNavbarComponent);
    })
});




describe('MainNavBarComponent', () => {
    
    it('expected a logged successful click on home', () => {
      cy.mount(MainNavbarComponent);
      cy.get('a').eq(0).click();
    });
  });
