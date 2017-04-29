import { InMemoryDbService } from 'angular-in-memory-web-api';
export class InMemoryDataServiceProduct implements InMemoryDbService {
  createDb() {
    let products = [

      {name: '普通分类1',price:11.11,image:'/assets/IMG_5097.JPG',buy:0,create_time:'2017-03-13T23:48:36+08:00',content:'产品1简介简介',id:1},
      {name: '普通分类1',price:22.22,image:'/assets/IMG_5098.JPG',buy:2,create_time:'2017-03-13T23:48:36+08:00',content:'产品1简介简介',id:2},
      {name: '普通分类1',price:33.33,image:'/assets/IMG_5099.JPG',buy:6,create_time:'2017-03-13T23:48:36+08:00',content:'产品1简介简介',id:3},
      {name: '普通分类1',price:44.44,image:'/assets/IMG_5100.JPG',buy:100,create_time:'2017-03-13T23:48:36+08:00',content:'产品1简介简介',id:4},

    ];
    return {products};
  }
}


/*
 Copyright 2016 Google Inc. All Rights Reserved.
 Use of this source code is governed by an MIT-style license that
 can be found in the LICENSE file at http://angular.io/license
 */
