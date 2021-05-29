import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
  signupForm: FormGroup;
  public path: string;
  constructor(
    private fb: FormBuilder,
    private http: HttpClient
  ) {
    this.path = 'placeholder';

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
    console.log(this.signupForm.value);
    this.http.post(this.path, this.signupForm.value).subscribe((result) =>{
        console.warn("Result: ", result)
    })
  }
}
