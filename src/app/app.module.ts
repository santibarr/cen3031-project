import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { purfectPartnerComponent } from './purfectPartner.component';
import { NavbarComponent } from './sharepage/navbar/navbar.component';
import { FooterComponent } from './sharepage/footer/footer.component';
import { HomeComponent } from './pages/home/home.component';
import { SheltersComponent } from './pages/shelters/shelters.component';
import { AboutComponent } from './pages/about/about.component';
import { ContactComponent } from './pages/contact/contact.component';
import { CatDetailsService } from './services/cat-details.service';


@NgModule({
  declarations: [
    AppComponent,
    purfectPartnerComponent,
    NavbarComponent,
    FooterComponent,
    HomeComponent,
    SheltersComponent,
    AboutComponent,
    ContactComponent,
    
    
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
  
  ],
  providers: [CatDetailsService],
  bootstrap: [AppComponent]
})
export class AppModule { }