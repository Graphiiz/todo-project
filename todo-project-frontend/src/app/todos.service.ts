import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { User } from './user.model';

@Injectable({
  providedIn: 'root'
})
export class TodosService {
  todos = ['eat','sleep','play'];
  name !: string;

  // change according to server ip
  backEndURL = 'http://localhost:1234/'

  constructor(private http: HttpClient) { }

  getLoginResponseFromEmail(email: string){
    // get name via http -> redirect to todos/:name
    console.log(this.backEndURL+'login/'+email);

    return this.http.get<User>(this.backEndURL+'login/'+email);

  }

  createNewUser(email: string){
    // create User object from email with empty array of todos
    const newUser: User = {
      email: email,
      todos: []
    }
    this.http.post<User>(this.backEndURL+'login/'+email, newUser).subscribe((response)=>{
      console.log(response)
    });
  }

  getTodoResponseFromEmail(email: string){

    return this.http.get<User>(this.backEndURL+'todos/'+email);

  }

  UpdateAndSaveToDB(userData: User){
    const email = userData["email"]
    this.http.post<User>(this.backEndURL+'todos/'+email, userData)
    .subscribe((response)=>{
      console.log(response)
    }, error => {
      console.log(error.message)
    });

  }

  deleteTodo(){
    // delete a single todo from todoList
    // In this project,we delete the data directly at the database. This function is not required.
  }

}
