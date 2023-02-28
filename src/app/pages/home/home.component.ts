import { Component, OnInit } from '@angular/core';
import { CatDetailsService } from 'src/app/services/cat-details.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit{
  constructor(private service: CatDetailsService) { }
  catData:any;
  ngOnInit():void{
    this.catData = this.service.catDetails;
  }
}

