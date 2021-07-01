import { Component, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { ProfileData, ProfiledataService } from '../profiledata.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  public profileJson: string = '';
  public usernameStore: string;
  public profileData: ProfileData;

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
        this.profileDataService.getProfileData(this.usernameStore).subscribe(
          res => {
            if (res === null) {
              this.profileDataService.createProfile(this.usernameStore)
              this.profileData = {
                id: 'NA',
                username: this.usernameStore,
                bio: '',
                upvotes: '',
                downvotes: ''
              }
              console.log(this.profileData)

            }
            else {
              this.profileData = res
              console.log(this.profileData)
            }
          }
        );
      }
    )
  }
  public setLocalStorage(): void {
    const data = JSON.parse(this.profileJson);
    this.usernameStore = data["preferred_username"]
    localStorage.setItem('username', data["preferred_username"]);
  }
}
