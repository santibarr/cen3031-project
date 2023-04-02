import { Component, OnInit } from '@angular/core';
import { CatDetailsService } from 'src/app/services/cat-details.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-cats',
  templateUrl: './cats.component.html',
  styleUrls: ['./cats.component.css']
})
export class CatsComponent implements OnInit {
  constructor(private http: HttpClient) {}
  catData: any;
  ngOnInit() {
    //Make an HTTP GET request to the API
    this.http.get('http://localhost:8000/cats').subscribe(data => {
      //Read the result field from the JSON response.
      this.catData = data;
      console.log(data)
    })
  }
}
