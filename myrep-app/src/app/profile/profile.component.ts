import { Component, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { ProfiledataService } from '../profiledata.service';

export interface ProfileData {
  username: string;
  bio: string;
}

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  public profileJson: string = '';
  private profileData: ProfileData;

  constructor(
    public auth: AuthService,
    public profileDataService: ProfiledataService
  ) {
  }

  ngOnInit(): void {
    this.auth.user$.subscribe(
      (profile) => {
        this.profileJson = JSON.stringify(profile, null, 2);
        this.setLocalStorage();
      }
    )

    this.profileData = this.profileDataService.getProfileData();
  }
  public setLocalStorage(): void {
    const data = JSON.parse(this.profileJson);
    localStorage.setItem('username', data["preferred_username"]);
  }
}