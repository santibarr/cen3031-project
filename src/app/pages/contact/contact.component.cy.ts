import { ContactComponent } from "./contact.component"

// This component test is for the home section of the website

describe('ContactComponent', () => {
    // mounts = visits but for components
    it('mounts', () => {
      cy.mount(ContactComponent)
      cy.contains('Contact');
    })
  })