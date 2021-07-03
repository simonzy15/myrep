import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProfileData, ProfiledataService } from '../profiledata.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  public targetUser: string;
  public profileData: ProfileData;
  public exists: boolean;

  constructor(
    private router: Router,
    public profileDataService: ProfiledataService
  ) { }

  ngOnInit(): void {
    this.targetUser = this.router.url.split('/')[2];
    this.profileDataService.getProfileData(this.targetUser).subscribe(
      res => {
        if (res === null) {
          this.exists = false
        }
        else {
          this.exists = true;
          this.profileData = res
        }
      }
    );
  }

}
