import { HomeComponent } from "./home.component"

// This component test is for the home section of the website

describe('HomeComponent', () => {
    // mounts = visits but for components
    it('mounts', () => {
      cy.mount(HomeComponent)
      cy.contains('Find Your Forever Feline Friend!');
    })
  })