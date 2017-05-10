import { Component, OnInit,ViewChild } from '@angular/core';
import { ProductService }         from '../Service/product.service';
import { CartService} from '../Service/cart.service'
import { Cart} from '../Models/Cart'
import { ActivatedRoute, Params } from '@angular/router';
import { Location }               from '@angular/common';
import { ModalDirective } from 'ngx-bootstrap/modal';

@Component({
  moduleId: module.id,
  selector: 'my-app',
  templateUrl: './cart.component.html',
  styleUrls: [ './cart.component.css']
})
export class CartComponent implements OnInit {

  carts: Cart[];
  count:number=0;
  public isShownMessage:boolean = false;
  public showMessage:string = "删除成功!";
  errorMessage: string;
  constructor(
    private productService: ProductService,
    private cartService: CartService,
    private route: ActivatedRoute,
    private location: Location
  ) { }

  deleteCar(id){
    this.route.params
      .switchMap((params: Params) => this.cartService.deleteCar(id))
      .subscribe(status => {
        if (status==10000){
          this.isShownMessage=true;
          setTimeout(() => {
            this.route.params
              .switchMap((params: Params) => this.cartService.getCar())
              .subscribe(
                carts => {
                  this.carts = carts;
                  this.count=0;
                  for(let cart of carts){
                    this.count+=(cart.Price*cart.Num);
                  }
                }
              );
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
  ngOnInit(): void {
    this.route.params
      .switchMap((params: Params) => this.cartService.getCar())
      .subscribe(
        carts => {
          this.carts = carts;
          this.count=0;
          for(let cart of carts){
            this.count+=(cart.Price*cart.Num);
          }
        }
      );
  }


}


