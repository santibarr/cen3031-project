import { Component, OnInit} from '@angular/core';
import { CatDetailsService } from 'src/app/services/cat-details.service';


@Component({
  selector: 'app-shelters',
  templateUrl: './shelters.component.html',
  styleUrls: ['./shelters.component.css']
})
export class SheltersComponent implements OnInit{
  constructor(private service: CatDetailsService) { }
  catData:any;
  
  ngOnInit(){
    this.catData = this.service.catDetails;
  }
  
}