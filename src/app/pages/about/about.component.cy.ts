import { AboutComponent } from "./about.component"

// This component test is for the home section of the website

describe('AboutComponent', () => {
    // mounts = visits but for components
    it('mounts', () => {
      cy.mount(AboutComponent)
      cy.contains('About');
    })
  })