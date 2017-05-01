import { Component, OnInit } from '@angular/core';
import { User }                from '../Models/User';
import { UserService }         from '../Service/user.service';
import { ActivatedRoute, Params } from '@angular/router';
import { Location }               from '@angular/common';

@Component({
  moduleId: module.id,
  selector: 'my-app',
  templateUrl: './user.component.html',
  styleUrls: [ './user.component.css']
})
export class UserComponent implements OnInit {
  user = new User();
  errorMessage: string;
  constructor(
    private userService: UserService,
    private route: ActivatedRoute,
    private location: Location
  ) { }

  ngOnInit(): void {

    this.userService.getUserInfo()
      .subscribe(userinfo => {
        this.user.Id = userinfo.Id;
        this.user.Name = userinfo.Name;
        this.user.Mobile = userinfo.Mobile;
        this.user.Account = userinfo.Account;
        this.user.Mark = userinfo.Mark;
        this.user.HeadImage = userinfo.HeadImage;
        this.user.Collection = userinfo.Collection;
        this.user.SelectAddress = userinfo.SelectAddress
      });
  }

}


