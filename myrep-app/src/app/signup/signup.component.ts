import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
  signupForm: FormGroup;
  
  constructor(
    private fb: FormBuilder
  ) {

  }

  ngOnInit(): void {
    this.initForm();
  }

  private initForm(): void {
    this.signupForm = this.fb.group({
      userName: '',
      password: '',
      email: ''
    })
  }
  onSubmit(): void{
    console.log(this.signupForm);
  }
}
