import { Injectable } from '@angular/core';
import { environment as env } from '../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

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
  constructor(
    private http: HttpClient
  ) {
    this.path = env.backendPath
  }
  
  public getProfileData(username: any): Observable<ProfileData> {
    return this.http.get<ProfileData>(
      this.path + '/api/getuser/' + username,
    )
  }

  public updateBio(username: string, newBio: string): Observable<any> {
    return this.http.put(
      this.path + '/api/edituser/' + username,
      JSON.stringify({
        'bio': newBio
      })
    )
  }
  
  public createProfile(username: any): void{
    if (username === null) {
      return
    }
    var body = JSON.stringify({
      'username': username
    })
    this.http.post<any>(
      this.path + '/api/register',
      body
    ).subscribe()
  }

}
