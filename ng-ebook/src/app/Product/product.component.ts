import { Component, OnInit } from '@angular/core';
import { Product }                from '../Models/Product';
import { ProductService }         from '../Service/product.service';
import { ActivatedRoute, Params } from '@angular/router';
import { Location }               from '@angular/common';
@Component({
  moduleId: module.id,
  selector: 'my-app',
  templateUrl: './product.component.html',
  styleUrls: [ './product.component.css' ]
})
export class ProductComponent implements OnInit {
  products: Product[];
  page:number = 1;
  hasPrev:boolean=false;
  hasNext:boolean=true;
  errorMessage: string;
  constructor(
    private productService: ProductService,
    private route: ActivatedRoute,
    private location: Location
  ) { }

  plus(){
    this.page++;
    this.route.params
      .switchMap((params: Params) => this.productService.pageProduct(+params['id'],this.page))
      .subscribe(products => this.products = products);
    this.route.params
      .switchMap((params: Params) => this.productService.pageHasNext(+params['id'],this.page))
      .subscribe(hasNext => this.hasNext = hasNext);
    if(this.page!=1){
      this.hasPrev = true
    }
    if(this.page==1){
      this.hasPrev = false
    }
  }

  minus(){
    if(this.page>1){
      this.hasPrev=true;
      this.page--;
      this.route.params
        .switchMap((params: Params) => this.productService.pageProduct(+params['id'],this.page))
        .subscribe(products => this.products = products);
      this.route.params
        .switchMap((params: Params) => this.productService.pageHasNext(+params['id'],this.page))
        .subscribe(hasNext => this.hasNext = hasNext);
      if(this.page!=1){
        this.hasPrev = true
      }
      if(this.page==1){
        this.hasPrev = false
      }
    }else{
      this.hasPrev=false;
    }
  }


  ngOnInit(): void {
    this.route.params
      .switchMap((params: Params) => this.productService.pageProduct(+params['id'],this.page))
      .subscribe(products => this.products = products);
    this.route.params
      .switchMap((params: Params) => this.productService.pageHasNext(+params['id'],this.page))
      .subscribe(hasNext => this.hasNext = hasNext);
  }

}


