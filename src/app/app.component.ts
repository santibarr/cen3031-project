import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'front-end';
  cats: any;

  constructor(private http: HttpClient) {}

  ngOnInit() {
    //Make an HTTP GET request to the API
    this.http.get('http://localhost:8080/cats').subscribe(data => {
      //Read the result field from the JSON response.
      this.cats = data;
      console.log(data)
    })

  }
}