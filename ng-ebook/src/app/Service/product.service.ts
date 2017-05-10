import { Injectable,Inject,OnInit }    from '@angular/core';
import { Headers, Http } from '@angular/http';
import 'rxjs/add/operator/toPromise';
import { Product } from '../Models/Product';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import {Response} from "@angular/http";
import {Observable} from 'rxjs/Observable';

import {ActivatedRoute, Params} from '@angular/router';
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
export class ProductService {
  private url = '/api/listProduct';
  private detailUrl = '/api/product';
  private hotUrl = '/api/hotProducts';
  private purl = '/api/product/listProduct';
  constructor(private http: Http,@Inject(ActivatedRoute) private router: ActivatedRoute) { }
  getProducts(id: number): Promise<Product[]> {
    const listUrl=`${this.url}?cid=${id}`;
    return this.http.get(listUrl)
      .toPromise()
      .then(response => {
        return response.json().products as Product[]
      })
      .catch(this.handleError);
  }
  getProduct(id: number): Promise<Product> {
    const newurl = `${this.detailUrl}?pid=${id}`;
    return this.http.get(newurl)
      .toPromise()
      .then(response => {
        return response.json().product as Product
      })
      .catch(this.handleError);
  }
  getHotProduct (): Observable<Product[]> {
    return this.http.get(this.hotUrl)
      .map(this.extractData)
      .catch(this.handleError);
  }

  pageProduct(id:number,page:number):Promise<Product[]>{
    const pageUrl=`${this.purl}/${id}?p=${page}`;
    return this.http.get(pageUrl)
      .toPromise()
      .then(response =>{
        return response.json().products as Product[]
      }).catch(this.handleError);
  }
  pageHasNext(id:number,page:number){
    const pageUrl=`${this.purl}/${id}?p=${page}`;
    return this.http.get(pageUrl)
      .toPromise()
      .then(response =>{
        return response.json().hasNext
      }).catch(this.handleError);
  }

  private extractData(res: Response) {
    let body = res.json();//解析为json
    // return body.categories || { };//解析body中自己包装的categories；Beego
    return body.products || { };//test
  }
  private extractData2(res: Response) {
    let body = res.json();//解析为json
    // return body.categories || { };//解析body中自己包装的categories；Beego
    console.error(body.product);

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
