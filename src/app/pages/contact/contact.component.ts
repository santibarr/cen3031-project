import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-contact',
  templateUrl: './contact.component.html',
  styleUrls: ['./contact.component.css']
})
export class ContactComponent implements OnInit {
  firstName: string | undefined;
  lastName: string | undefined;
  email: string | undefined;
  phone: string | undefined;
  
  constructor(private http: HttpClient) { }
  ngOnInit(): void {
    //print debug
    //console.log("created");
  }

  onFormSubmit(event: Event) {
    
    //send the user data to the server
    const User = {
      firstName: this.firstName,
      lastName: this.lastName,
      email: this.email,
      phone: this.phone
    };

    //print debug
    //console.log("helloooooo")

    this.http.post('http://localhost:8000/contact', User).subscribe(data => {
      console.log(data);
      
    });
  }
}
