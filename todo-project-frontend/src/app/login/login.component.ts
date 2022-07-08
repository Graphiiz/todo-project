import { HttpClient } from '@angular/common/http';
import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, NgForm, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { flatMap, map } from 'rxjs/operators';
import { TodosService } from '../todos.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm !: FormGroup;
  placeholder = 'example@mail.com'
  email !: string;
  name = 'default';

  constructor(private router: Router, private route: ActivatedRoute,
    private http: HttpClient, private todoService: TodosService) { }

  ngOnInit(): void {
    this.loginForm = new FormGroup({
      'email': new FormControl(null, [Validators.required, Validators.email])
    })
  }

  onSubmit(){
    console.log('Submitted!');
    console.log(this.loginForm.status);
    console.log(this.loginForm.get('email')?.valid)
    // this.loginForm.reset(); // reset form after submitted, may not be neccessary if page switch after submit
    if(this.loginForm.get('email')?.valid){
      // if email is valid [actually valid because the form can be submitted if email is valid :)]
      this.email = this.loginForm.get('email')?.value
      console.log(this.email)

      // retrieve name from db -> will get name if the user is exist or when new user is created completely
      // since the component use the value from http response -> subscribe in component not the service.
      this.todoService.getLoginResponseFromEmail(this.email)
      .pipe(map(response => {
        const responseString = JSON.stringify(response); // object -> string
        const responseJSON = JSON.parse(responseString); // string -> json
        console.log(responseJSON)
        return responseJSON;
      }))
      .subscribe(userData =>{
        const email = userData["email"];
        // redirect to the url todos/:email
        this.router.navigateByUrl('todos/'+email)
      }, error => {
        console.log(error)
        this.todoService.createNewUser(this.email)
        alert('User is created.')
      });
    } else {
      this.router.navigate(['../'], {relativeTo: this.route});
      console.log('Back to login page.')
    }
    // this.loginForm.reset(); // reset form after submitted, may not be neccessary if page switch after submit
  }

}

