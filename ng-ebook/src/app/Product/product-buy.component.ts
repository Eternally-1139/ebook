import { Component, OnInit,ViewChild } from '@angular/core';
import { Product }                from '../Models/Product';
import { ProductService }         from '../Service/product.service';
import { ActivatedRoute, Params } from '@angular/router';
import { Location }               from '@angular/common';
import { ModalDirective } from 'ngx-bootstrap/modal';
@Component({
  moduleId: module.id,
  selector: 'my-app',
  templateUrl: './product-buy.component.html',
  styleUrls: [ './product-buy.component.css' ]
})
export class ProductBuyComponent implements OnInit {
  isDisplay:"none";
  modelText = "";
  products = new Product();
  errorMessage: string;
  constructor(
    private productService: ProductService,
    private route: ActivatedRoute,
    private location: Location
  ) { }

  @ViewChild('autoShownModal') public autoShownModal:ModalDirective;
  public isModalShown:boolean = false;

  public showModal():void {
    this.isModalShown = true;
  }

  public hideModal():void {
    this.autoShownModal.hide();
  }

  public onHidden():void {
    this.isModalShown = false;
  }
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


