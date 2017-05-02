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
  errorMessage: string;
  constructor(
    private productService: ProductService,
    private cartService: CartService,
    private route: ActivatedRoute,
    private location: Location
  ) { }
  ngOnInit(): void {
    this.route.params
      .switchMap((params: Params) => this.cartService.getCar())
      .subscribe(carts => this.carts = carts);
  }


}


