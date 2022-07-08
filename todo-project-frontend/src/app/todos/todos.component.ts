import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { TodosService } from '../todos.service';

@Component({
  selector: 'app-todos',
  templateUrl: './todos.component.html',
  styleUrls: ['./todos.component.css'],
  providers: [TodosService]
})
export class TodosComponent implements OnInit {
  todoForm !: FormGroup;
  placeholder = "todo";
  todoList !: string[];
  constructor(private todosService: TodosService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    // initialize the form
    this.todoForm = new FormGroup({
      'todo': new FormControl(null, Validators.required)
    })
    // retrieve the id from route parameter
    const email = this.route.snapshot.params["email"]
    this.todosService.getTodoResponseFromEmail(email)
    .subscribe(response => {
      this.todoList = response["todos"];
    })

    // initialize the todoList from service
    // this.todoList = this.todosService.getTodos()
    // get from db via http
  }

  onSubmit(){
    console.log(this.todoForm);
    if (!this.todoForm.get('todo')?.invalid){
      this.todoList.push(this.todoForm.get('todo')?.value);
    }
    this.todoForm.reset();
  }

  onDeleteTodo(i:number){
    this.todoList = this.todoList.filter(todo => this.todoList.indexOf(todo)!== i)
  }

  onSave(){
    const email = this.route.snapshot.params["email"]
    const todos = this.todoList
    this.todosService.UpdateAndSaveToDB({email:email,todos:todos});
  }

}
