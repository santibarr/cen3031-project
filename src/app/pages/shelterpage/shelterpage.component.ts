import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CatDetailsService } from 'src/app/services/cat-details.service';

@Component({
  selector: 'app-shelterpage',
  templateUrl: './shelterpage.component.html',
  styleUrls: ['./shelterpage.component.css']
})
export class ShelterpageComponent implements OnInit {
  constructor(private param:ActivatedRoute,private service:CatDetailsService) { }
  getcatsId:any;
  catsData:any;
  ngOnInit(): void {
    this.getcatsId = this.param.snapshot.paramMap.get('id');
    console.log(this.getcatsId,'getcats');
    if(this.getcatsId)
    {
      this.catsData = this.service.catDetails.filter((value)=>{
        return value.id == this.getcatsId;
      });
      console.log(this.catsData,'shetlersdata>>');
      
    }
    
  }
}
