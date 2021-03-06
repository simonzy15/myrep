import { Component, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  constructor(
    public auth: AuthService,
  ) {
  }

  ngOnInit(): void {
  }

  public login(): void {
    this.auth.loginWithRedirect({ 
      redirect_uri: window.location.origin,
      appState: { target: '/profile'}
    })
  }
  public signup(): void {
    this.auth.loginWithRedirect({
      screen_hint: 'signup',
      redirect_uri: window.location.origin,
      appState: { target: '/profile'}
    })
  }
}
