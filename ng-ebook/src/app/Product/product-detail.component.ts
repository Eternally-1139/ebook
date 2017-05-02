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

  public addCart():void{
    this.cartService.addCar(this.products.Id,this.products.Name,this.products.Price,this.products.Image,1,this.products.Content)
      .subscribe(status =>{
          if (status==10000){
            this.add()
          }else{
            this.add()
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


