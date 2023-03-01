import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CatDetailsService } from 'src/app/services/cat-details.service';
import { __values } from 'tslib';

@Component({
  selector: 'app-adoptionpage',
  templateUrl: './adoptionpage.component.html',
  styleUrls: ['./adoptionpage.component.css']
})
export class AdoptionpageComponent implements OnInit {
  constructor(private param:ActivatedRoute, private service:CatDetailsService) { }
  getShelterId:any;
  shelterData:any;
  ngOnInit(): void {
    this.getShelterId = this.param.snapshot.paramMap.get('id');
    console.log(this.getShelterId, 'getshelter');
    if(this.getShelterId) {
      this.shelterData = this.service.catDetails.filter((value)=>{
        return value.id ==  this.getShelterId;
      });
      console.log(this.shelterData, 'shelterdata>>');
    }
  }
}
