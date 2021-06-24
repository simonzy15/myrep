import { Injectable } from '@angular/core';
import { environment as env } from '../environments/environment';
import { HttpClient } from '@angular/common/http';
import { ProfileData } from './profile/profile.component';

@Injectable({
  providedIn: 'root'
})
export class ProfiledataService {
  public state: string;
  public path: string;
  constructor(
    private http: HttpClient
  ) {
    this.path = env.backendPath
  }

  public getProfileData(): ProfileData {
    this.http.get(
      this.path + '/api/getuser/12'
    ).subscribe(
      res => {
        console.log(res)
      }
    )
    return {
      username: '',
      bio: ''
    }
  }
}
