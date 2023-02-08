import { Component } from '@angular/core';

@Component({
    selector: 'app-purfectPartner',
    template: '<h1>{{title}}</h1>',
    styles: [`
    h1 {
        color: pink;
    }`]
})

export class purfectPartnerComponent {
    title = 'Purfect Partner!';
}