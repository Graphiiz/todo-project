import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { TodosComponent } from './todos/todos.component';


const routes: Routes = [
  {path: '', redirectTo: '/login', pathMatch: 'full'}, //redirect if fullpath is empty
  {path: 'login', component: LoginComponent},
  {path: 'todos/:email', component: TodosComponent}
]
// multipaths can point to a component.

@NgModule({
  declarations: [],
  imports: [
    RouterModule.forRoot(routes)
  ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }
