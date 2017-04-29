import { InMemoryDbService } from 'angular-in-memory-web-api';
export class InMemoryDataService implements InMemoryDbService {
  createDb() {
    let categories = [
      {Id: 1, Name: '普通分类1',Status:0,Image:''},
      {Id: 2, Name: '普通分类2',Status:0,Image:''},
      {Id: 3, Name: '普通分类3',Status:0,Image:''},
      {Id: 4, Name: '普通分类4',Status:0,Image:''},
      {Id: 5, Name: '普通分类5',Status:0,Image:''},
      {Id: 6, Name: '普通分类6',Status:0,Image:''},
      {Id: 7, Name: '普通分类7',Status:0,Image:''},
      {Id: 8, Name: '普通分类8',Status:0,Image:''}
    ];
    return {categories};
  }
}


/*
 Copyright 2016 Google Inc. All Rights Reserved.
 Use of this source code is governed by an MIT-style license that
 can be found in the LICENSE file at http://angular.io/license
 */
