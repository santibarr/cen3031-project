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
  phoneNumber: string | undefined;
  
  constructor(private http: HttpClient) { }
  ngOnInit(): void {
    console.log("created");
  }

  onFormSubmit(event: Event) {
    
    //send the user data to the server
    const User = {
      firstName: this.firstName,
      lastName: this.lastName,
      email: this.email,
      phoneNumber: this.phoneNumber
    };

    console.log("helloooooo")

    this.http.post('http://localhost:8000/contact', User).subscribe(data => {
      console.log(data);
    });
  }
}
