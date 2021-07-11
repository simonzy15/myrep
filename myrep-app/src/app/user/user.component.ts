import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProfileData, ProfiledataService } from '../profiledata.service';
import { FormGroup, FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  commentForm: FormGroup;
  public targetUser: string;
  public profileData: ProfileData;
  public exists: boolean;

  constructor(
    private fb: FormBuilder,
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
          this.initForm()
        }
      }
    );
  }

  private initForm(): void {
    this.commentForm = this.fb.group({
      comment: ''
    })
  }
  onSubmit(): void {
    const newComment = this.commentForm.value['comment']
    console.log(newComment)
    this.commentForm.reset()
  }
}
