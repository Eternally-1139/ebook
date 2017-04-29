import { Component, OnInit } from '@angular/core';
import { Category }                from '../Models/Category';
import { Product }                from '../Models/Product';
import { Swipe }                from '../Models/Swipe';
import { CategoryService }         from '../Service/category.service';
import { ProductService }         from '../Service/product.service';
import { SwipeService }         from '../Service/swipe.service';
import { ActivatedRoute, Params } from '@angular/router';


@Component({
  moduleId: module.id,
  selector: 'my-app',
  templateUrl: './home.component.html',
  styleUrls: [ './home.component.css' ]
})
export class HomeComponent implements OnInit {
  products: Product[];
  categories: Category[];
  swipe: Swipe[];
  errorMessage: string;
  constructor(
    private categoryService: CategoryService,
    private  productService: ProductService,
    private route: ActivatedRoute,
    private swipeService:SwipeService,
  ) { }
  getCategories() {
    this.categoryService.getCategories()
      .subscribe(
        categories => this.categories = categories,
        error =>  this.errorMessage = <any>error);
  }
  getHotProduct() {
    this.productService.getHotProduct()
      .subscribe(products => this.products = products);
  }
  getSwipe(){
    this.swipeService.getSwipe()
      .subscribe(
        swipe => this.swipe = swipe);
  }

  ngOnInit(): void {
    this.getCategories();
    this.getHotProduct();
    this.getSwipe();

  }


}
