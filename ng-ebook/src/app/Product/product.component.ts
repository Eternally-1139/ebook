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
  nextDisabled:boolean=false;
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
  }


  ngOnInit(): void {
    this.route.params
      .switchMap((params: Params) => this.productService.pageProduct(+params['id'],this.page))
      .subscribe(products => this.products = products);
  }

}


