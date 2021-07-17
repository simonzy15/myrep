import { Component, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { ProfiledataService } from 'src/app/profiledata.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  public userPage: string
  constructor(
    public auth: AuthService,
    public profileDataService: ProfiledataService
  ) {
    
  }

  ngOnInit(): void {
    this.profileDataService.subject.subscribe(
      res => {
        this.userPage = "/users/" + res
      }
    )
  }
}
