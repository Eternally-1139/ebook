import { Injectable,Inject,OnInit }    from '@angular/core';
import { Headers, Http } from '@angular/http';
import 'rxjs/add/operator/toPromise';
import { Swipe } from '../Models/Swipe';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import {Response} from "@angular/http";
import {Observable} from 'rxjs/Observable';
// Observable class extensions
import 'rxjs/add/observable/of';
import 'rxjs/add/observable/throw';

// Observable operators
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/debounceTime';
import 'rxjs/add/operator/distinctUntilChanged';
import 'rxjs/add/operator/do';
import 'rxjs/add/operator/filter';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/switchMap';

@Injectable()
export class SwipeService {
  private url = '/api/listCarousel';
  constructor(private http: Http) { }

  getSwipe (): Observable<Swipe[]> {
    return this.http.get(this.url)
      .map(this.extractData)
      .catch(this.handleError);
  }

  private extractData(res: Response) {
    let body = res.json();//解析为json
    // return body.categories || { };//解析body中自己包装的categories；Beego
    return body.swipe || { };//test
  }
  private handleError (error: Response | any) {
    // In a real world app, we might use a remote logging infrastructure
    let errMsg: string;
    if (error instanceof Response) {
      const body = error.json() || '';
      const err = body.error || JSON.stringify(body);
      errMsg = `${error.status} - ${error.statusText || ''} ${err}`;
    } else {
      errMsg = error.message ? error.message : error.toString();
    }
    console.error(errMsg);
    return Observable.throw(errMsg);
  }
  ngOnInit() {

  }

}
