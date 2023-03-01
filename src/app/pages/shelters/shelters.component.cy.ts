import { SheltersComponent } from "./shelters.component"

// This component test is for the home section of the website

describe('SheltersComponent', () => {
    // mounts = visits but for components
    it('mounts', () => {
      cy.mount(SheltersComponent)
      cy.contains('shelters');
    })
  })