import { Component, OnInit,ViewChild } from '@angular/core';
import { Product }                from '../Models/Product';
import { ProductService }         from '../Service/product.service';
import { CartService} from '../Service/cart.service'
import { Cart} from '../Models/Cart'
import { ActivatedRoute, Params } from '@angular/router';
import { Location }               from '@angular/common';
import { ModalDirective } from 'ngx-bootstrap/modal';
import {Subject} from 'rxjs/Subject';

@Component({
  moduleId: module.id,
  selector: 'my-app',
  templateUrl: './product-detail.component.html',
  styleUrls: [ './product-detail.component.css']
})
export class ProductDetailComponent implements OnInit {

  private _success = new Subject<string>();
  staticAlertClosed = false;
  successMessage: string;
  isDisplay:string="none";
  modelText = "";
  @ViewChild('autoShownModal') public autoShownModal:ModalDirective;
  public isModalShown:boolean = false;

  public showModal():void {
    this.isModalShown = true;
    setTimeout(() => {
      this.isModalShown = false
    }, 2000);
  }

  public hideModal():void {
    this.autoShownModal.hide();
  }

  public onHidden():void {
    this.isModalShown = false;
  }


  products = new Product();
  errorMessage: string;
  constructor(
    private productService: ProductService,
    private cartService: CartService,
    private route: ActivatedRoute,
    private location: Location
  ) { }
  public changeSuccessMessage() {
    this._success.next(`添加购物车成功`);
  }

  public addCart():void{
    this.cartService.addCar(this.products.Id,this.products.Name,this.products.Price,this.products.Image,1,this.products.Content)
      .subscribe(status =>{
          if (status==10000){
            this.changeSuccessMessage()
          }else{
            this.changeSuccessMessage()
          }
        });
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

    setTimeout(() => this.staticAlertClosed = true, 20000);

    this._success.subscribe((message) => this.successMessage = message);
    this._success.debounceTime(5000).subscribe(() => this.successMessage = null);

  }


}


