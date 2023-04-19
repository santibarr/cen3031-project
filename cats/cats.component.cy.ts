import { CatsComponent } from "./cats.component"

// This component test is for the home section of the website

describe('CatsComponent', () => {
    // mounts = visits but for components
    it('mounts', () => {
      cy.mount(CatsComponent)
      cy.contains('Slay');
    })
  })

  describe('Adopt Button', () => {
    // mounts = visits but for button
    it('clicks', () => {
      cy.mount(CatsComponent)
      cy.contains('Adopt Now').click;
    })
  })