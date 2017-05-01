import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import {RouterModule, Routes} from "@angular/router";
import { CarouselModule } from 'ngx-bootstrap';
import { AppComponent } from './app.component';
import { requestOptionsProvider }   from './default-request-options.service';
//导入各页面组件
import { HomeComponent} from './Home/home.component'
import { UserComponent} from './User/user.component'
import { CartComponent} from './Cart/cart.component'
import { MessageComponent} from './Message/message.component'
import { ProductComponent} from './Product/product.component'
import { ProductDetailComponent} from './Product/product-detail.component'
import { ProductBuyComponent} from './Product/product-buy.component'
//导入service
import {CategoryService} from './Service/category.service'
import {ProductService} from './Service/product.service'
import {SwipeService} from './Service/swipe.service'
import {UserService} from './Service/user.service'

//前端模拟接口api
import { InMemoryDataService }  from './in-memory-data.service';//test
import { InMemoryDataServiceProduct }  from './in-memory-data-product.service';//test
import { InMemoryWebApiModule } from 'angular-in-memory-web-api';//test

const appRoutes: Routes = [
  {
    path: '',
    component: HomeComponent,
  },
  { path: 'user',
    component: UserComponent,
  },
  { path: 'message',
    component: MessageComponent,
  },
  { path: 'cart',
    component: CartComponent,
  },
  //子组件
  { path: 'product/:id',
    component: ProductComponent
  },
  { path: 'detail/:id',
    component: ProductDetailComponent
  },
  { path: 'buy/:id',
    component: ProductBuyComponent
  },
];

@NgModule({
  declarations: [
    AppComponent,HomeComponent,UserComponent,CartComponent,MessageComponent,ProductComponent,ProductDetailComponent,ProductBuyComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    RouterModule.forRoot(appRoutes),
    CarouselModule.forRoot(),
    // InMemoryWebApiModule.forRoot(InMemoryDataService),//test
    // InMemoryWebApiModule.forRoot(InMemoryDataServiceProduct),//test
    // InMemoryWebApiModule.forRoot(InMemoryDataService),

  ],
  providers: [CategoryService,requestOptionsProvider,ProductService,SwipeService,UserService],
  bootstrap: [AppComponent]
})
export class AppModule { }
