import { Component, OnInit } from '@angular/core';
import { ProfiledataService } from '../profiledata.service';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  constructor(
    private profileDataService: ProfiledataService
  ) {
  }

  ngOnInit(): void {
    var userName = localStorage.getItem('username')
    if (userName !== null) {
      this.profileDataService.currentUser.next(userName)
    }
  }
}
