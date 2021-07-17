import { Injectable } from '@angular/core';
import { environment as env } from '../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';

export interface ProfileData {
  id: string;
  username: string;
  bio: string;
  upvotes: string;
  downvotes: string;
  picture: string;
}

@Injectable({
  providedIn: 'root'
})
export class ProfiledataService {
  public state: string;
  public path: string;
  public subject = new Subject<string>();
  constructor(
    private http: HttpClient
  ) {
    this.path = env.backendPath
  }
  
  public getProfileData(username: any): Observable<ProfileData> {
    this.subject.next(username)
    return this.http.get<ProfileData>(
      this.path + '/api/getuser/' + username,
    )
  }

  public updateBio(username: string, newBio: string): Observable<any> {
    var body = JSON.stringify({
      'bio': newBio
    })
    return this.http.put(
      this.path + '/api/edituser/' + username,
      body
    )
  }
  
  public createProfile(username: any, picture: string): void{
    if (username === null) {
      return
    }
    var body = JSON.stringify({
      'username': username,
      'picture': picture
    })
    this.http.post<any>(
      this.path + '/api/register',
      body
    ).subscribe()
  }

}
