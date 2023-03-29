import { Component, OnInit } from '@angular/core';
import { CatDetailsService } from 'src/app/services/cat-details.service';

@Component({
  selector: 'app-cats',
  templateUrl: './cats.component.html',
  styleUrls: ['./cats.component.css']
})
export class CatsComponent implements OnInit {
  constructor(private service:CatDetailsService) { }
  catData:any;
  ngOnInit(): void {
    this.catData = this.service.catDetails;
  }
}
