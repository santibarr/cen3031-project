import { CatsComponent } from "./cats.component"

// This component test is for the home section of the website

describe('CatsComponent', () => {
    // mounts = visits but for components
    it('mounts', () => {
      cy.mount(CatsComponent)
      cy.contains('Cats');
    })
  })