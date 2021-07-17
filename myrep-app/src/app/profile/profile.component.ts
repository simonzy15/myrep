import { Component, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { ProfileData, ProfiledataService } from '../profiledata.service';
import { FormGroup, FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  profileForm: FormGroup;
  public profileJson: string = '';
  public usernameStore: string;
  public profileData: ProfileData;

  constructor(
    private fb: FormBuilder,
    public auth: AuthService,
    public profileDataService: ProfiledataService
  ) {
  }

  ngOnInit(): void {
    this.auth.user$.subscribe(
      (profile) => {
        this.profileJson = JSON.stringify(profile, null, 2);
        const data = JSON.parse(this.profileJson)
        this.setLocalStorage();
        this.profileDataService.getProfileData(this.usernameStore).subscribe(
          res => {
            if (res === null) {
              this.profileDataService.createProfile(this.usernameStore, data["picture"])
              this.profileData = {
                id: 'NA',
                username: this.usernameStore,
                bio: '',
                upvotes: '',
                downvotes: '',
                picture: data["picture"]
              }

            }
            else {
              this.profileData = res
            }
            this.profileDataService.currentUser.next(this.usernameStore)
            this.initForm();
          }
        );
      }
    )
  }

  private initForm(): void {
    this.profileForm = this.fb.group({
      bio: this.profileData.bio
    })
  }

  public setLocalStorage(): void {
    const data = JSON.parse(this.profileJson);
    this.usernameStore = data["preferred_username"]
    localStorage.setItem('username', data["preferred_username"]);
  }

  onSubmit(): void {
    const newBio = this.profileForm.value['bio']
    if (newBio !== this.profileData.bio) {
      this.profileDataService.updateBio(this.profileData.username, newBio).subscribe();
    }
  }
}
