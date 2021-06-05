import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm: FormGroup;
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
    this.loginForm = this.fb.group({
      Username: '',
      Password: ''
    })
  }
  onSubmit(): void{
    console.log(this.loginForm);
    console.log(this.loginForm.value);
    this.http.post(this.path, this.loginForm.value).subscribe((result) =>{
        console.warn("Result: ", result)
    })
  }
}
