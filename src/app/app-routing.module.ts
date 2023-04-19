import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutComponent } from './pages/about/about.component';
import { ContactComponent } from './pages/contact/contact.component';
import { HomeComponent } from './pages/home/home.component';
import { ShelterpageComponent } from './pages/shelterpage/shelterpage.component';
import { CatsComponent } from './pages/cats/cats.component';
import { QuizComponent } from './pages/quiz/quiz.component';

const routes: Routes = [
  {path:'',component:HomeComponent},
  {path:'cats',component:CatsComponent},
  {path:'cats/:id',component:ShelterpageComponent},
  {path:'about',component:AboutComponent},
  {path:'contact',component:ContactComponent},
  {path:'quiz',component:QuizComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
