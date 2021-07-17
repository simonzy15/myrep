import { Component, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  public userPage: string
  constructor(
    public auth: AuthService
  ) {
    
  }

  ngOnInit(): void {
    this.userPage = "/users/" + localStorage.getItem('username')
  }
}
