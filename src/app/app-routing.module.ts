import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginPageComponent } from './login-page/login-page.component';
import { SignupPageComponent } from './signup-page/signup-page.component';
import { MainNavbarComponent } from './main-navbar/main-navbar.component';
import { HomePageComponent } from './home-page/home-page.component';
import { AboutPageComponent } from './about-page/about-page.component';
import { TextboxPageComponent } from './textbox-page/textbox-page.component';
import { HistoryPageComponent } from './history-page/history-page.component';
import { EntryPageComponent } from './entry-page/entry-page.component'
import { MoodChartComponent } from './mood-chart/mood-chart.component';
import { StatsPageComponent } from './stats-page/stats-page.component';


export const routes: Routes = [
  {path: 'home', component:HomePageComponent},
  {path:'login', component:LoginPageComponent},
  {path:'about', component:AboutPageComponent},
  {path:'sign-up', component:SignupPageComponent},
  {path:'journal', component:TextboxPageComponent},
  {path: 'history', component:HistoryPageComponent},
  {path:'entry', component:EntryPageComponent},
  {path: 'mood', component:MoodChartComponent},
  {path:'stats',component:StatsPageComponent},
  {path:'', redirectTo:'/home', pathMatch:'full'},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
