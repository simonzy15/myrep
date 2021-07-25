import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '@auth0/auth0-angular';
import { ProfiledataService } from 'src/app/profiledata.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  public userPage: string
  public currUrl: string
  constructor(
    private router: Router,
    public auth: AuthService,
    public profileDataService: ProfiledataService
  ) {
  }

  ngOnInit(): void {
    this.currUrl = this.router.url
    var userName = localStorage.getItem('username')
    if (userName !== null) {
      this.userPage = "/users/" + userName
    } else {
      this.profileDataService.currentUser.subscribe(
      res => {
        this.userPage = "/users/" + res
      }
    )
    }
  }
  public toProfile(): void {
    if(this.router.url.includes('/users/') && this.router.url != this.userPage){
      this.router.navigateByUrl(this.userPage).then(() => {
        window.location.reload()
      })
    }
    else {
      this.router.navigateByUrl(this.userPage)
    }
  }
}
