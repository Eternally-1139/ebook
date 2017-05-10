import { Component, OnInit,ViewChild } from '@angular/core';
import { Product }                from '../Models/Product';
import { ProductService }         from '../Service/product.service';
import { CartService} from '../Service/cart.service'
import { Cart} from '../Models/Cart'
import { ActivatedRoute, Params } from '@angular/router';
import { Location }               from '@angular/common';
import { ModalDirective } from 'ngx-bootstrap/modal';

@Component({
  moduleId: module.id,
  selector: 'my-app',
  templateUrl: './product-detail.component.html',
  styleUrls: [ './product-detail.component.css']
})
export class ProductDetailComponent implements OnInit {
  isDisplay:string="none";
  public isShownMessage:boolean = false;
  public showMessage:string = "亲，我在购物车等你哦！"
  modelText = "";
  @ViewChild('autoShownModal') public autoShownModal:ModalDirective;


  products = new Product();
  errorMessage: string;
  constructor(
    private productService: ProductService,
    private cartService: CartService,
    private route: ActivatedRoute,
    private location: Location
  ) { }

  public addCart():void{
    this.cartService.addCar(this.products.Id,this.products.Name,this.products.Price,this.products.Image,1,this.products.Content)
      .subscribe(status =>{
        if (status==10000){
          this.isShownMessage=true;
          setTimeout(() => {
            this.isShownMessage=false;
          }, 1500);
        }else{
          this.isShownMessage=true;
          this.showMessage="好像出错了，请检查您的网络";
          setTimeout(() => {
            this.isShownMessage=false;
          }, 1500);
        }
      });
  }

  public alerts: any = [];
  public add(): void {
    this.alerts.push({
      type: 'info',
      msg: `加入购物车成功`,
      timeout: 1500
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

  }

}


