import { Component } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { User, UserService } from "openapi"
import { Subscriber } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  getUserForm = new FormGroup({
    id: new FormControl('1')
  });
  
  createUserForm = new FormGroup({
    id: new FormControl('2'),
    nickname: new FormControl('create-test')
  });

  updateUserForm = new FormGroup({
    id: new FormControl('3'),
    nickname: new FormControl('update-test')
  });

  deleteUserForm = new FormGroup({
    id: new FormControl('4')
  });

  constructor(private readonly userService: UserService) { }

  public getUser(): void {
    this.userService
      .getUser(this.getUserForm.value.id)
      .subscribe((user) => {
        this.updateUserForm.patchValue(user);
        console.log('Get user: ' + user);
      });
  }

  public createUser(): void {
    let createUser: User = {
      id : this.createUserForm.value.id,
      nickname: this.createUserForm.value.nickname 
    };

    this.userService
      .createUser(createUser)
      .subscribe((user) => {
        console.log('Create user: ' + user);
      });
  }

  public updateUser(): void {
    let createUser: User = {
      id : this.updateUserForm.value.id,
      nickname: this.updateUserForm.value.nickname 
    };

    this.userService
      .updateUser(createUser)
      .subscribe((user) => {
        console.log('Update user: ' + user);
      });
  }

  public deleteUser(): void {
    this.userService
      .deleteUser(this.deleteUserForm.value.id)
      .subscribe((user) => {
        console.log('Delete user: ' + user);
      });
  }
}
