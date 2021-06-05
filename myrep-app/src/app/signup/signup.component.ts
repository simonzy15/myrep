import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';

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
    this.path = 'http://localhost:8001/api/register';

  }

  ngOnInit(): void {
    this.initForm();
  }

  private initForm(): void {
    this.signupForm = this.fb.group({
      Username: '',
      Password: '',
      Email: ''
    })
  }
  onSubmit(): void{
    console.log(this.signupForm.value);
    
    const httpHeader = new HttpHeaders({
      'content-type' : 'application/x-www-form-urlencoded'
    })
    this.http.post(
      this.path, 
      this.signupForm.value,
      {
        headers: httpHeader
      }
      ).subscribe((result) =>{
        console.warn("Result: ", result)
    })
  }
}
