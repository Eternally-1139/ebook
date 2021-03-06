import { Component, OnInit } from '@angular/core';
import { Product }                from '../Models/Product';
import { ProductService }         from '../Service/product.service';
import { ActivatedRoute, Params } from '@angular/router';
import { Location }               from '@angular/common';

@Component({
  moduleId: module.id,
  selector: 'my-app',
  templateUrl: './product-rent.component.html',
  styleUrls: [ './product-rent.component.css' ]
})
export class ProductRentComponent implements OnInit {
  products = new Product();
  errorMessage: string;
  constructor(
    private productService: ProductService,
    private route: ActivatedRoute,
    private location: Location
  ) { }

  locationBack(): void {
    this.location.back();
  }
  ngOnInit(): void {
    this.route.params
      .switchMap((params: Params) => this.productService.getProduct(+params['id']))
      .subscribe(product => {
        this.products.Name = product.Name;
        this.products.Content = product.Content;
        this.products.Price = product.Price;
        this.products.Buy = product.Buy;
        this.products.Id = product.Id;
        this.products.Image = product.Image;
        this.products.Pic = product.Pic;
      });
  }

}


