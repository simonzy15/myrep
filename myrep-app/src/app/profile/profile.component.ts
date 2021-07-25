import { Component, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { ProfileData, ProfiledataService } from '../profiledata.service';
import { FormGroup, FormBuilder } from '@angular/forms';
import { MatDialog } from "@angular/material/dialog";
import { DialogComponent } from '../component/dialog/dialog.component';
import { HttpClient } from '@angular/common/http';
import { environment as env } from '../../environments/environment';

export interface EditPicture {
  picture: string;
}

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
  public changePicture: string;
  public path: string;

  constructor(
    private fb: FormBuilder,
    private http: HttpClient,
    public auth: AuthService,
    public profileDataService: ProfiledataService,
    public dialog: MatDialog
  ) {
    this.path = env.backendPath
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
                upvotes: '0',
                downvotes: '0',
                picture: data["picture"]
              }

            }
            else {
              this.profileData = res
              localStorage.setItem('picture', this.profileData.picture)
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

  public openDialog() {
    const dialogRef = this.dialog.open(DialogComponent, {
      data: {
        picture: this.changePicture
      }
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result != undefined && result != ''){
        this.changePicture = result
        this.profileData.picture = this.changePicture
        var body = JSON.stringify({
          'username': this.profileData.username,
          'picture': this.changePicture,
        })
        this.http.put(
          this.path + '/api/updatephoto',
          body
        ).subscribe()
      }
    });
  }

  public setLocalStorage(): void {
    const data = JSON.parse(this.profileJson);
    this.usernameStore = data["preferred_username"]
    localStorage.setItem('username', data["preferred_username"]);
    localStorage.setItem('picture', data["picture"])
  }

  onSubmit(): void {
    const newBio = this.profileForm.value['bio']
    if (newBio !== this.profileData.bio) {
      this.profileDataService.updateBio(this.profileData.username, newBio).subscribe();
    }
  }
}
