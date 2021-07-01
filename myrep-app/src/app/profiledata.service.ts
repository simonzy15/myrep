import { Injectable } from '@angular/core';
import { environment as env } from '../environments/environment';
import { HttpClient, HttpHeaders } from '@angular/common/http';

export interface ProfileData {
  id: string;
  username: string;
  bio: string;
  upvotes: string;
  downvotes: string;
}

@Injectable({
  providedIn: 'root'
})
export class ProfiledataService {
  public state: string;
  public path: string;
  public profileData: ProfileData;
  public usernameStore: any;
  constructor(
    private http: HttpClient
  ) {
    this.path = env.backendPath
  }
  
  public getProfileData(): void {
    this.usernameStore = localStorage.getItem('username');
    this.http.get<ProfileData>(
      this.path + '/api/getuser/testaccount',
    ).subscribe(
      res => {
        console.log(res)
        if (res === null) {
          this.createProfile(this.usernameStore)
        }
        else {
          this.profileData = res
        }
      }
    )
  }
  
  public createProfile(username: any): void{
    if (username === null) {
      return
    }
    var body = JSON.stringify({
      'username': username
    })
    console.log(body)
    this.http.post<any>(
      this.path + '/api/register',
      body
    ).subscribe(
      res => {
        console.log(res)
      }
    )
  }

}
