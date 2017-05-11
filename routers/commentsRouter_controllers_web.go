package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "ListCarousel",
			Router: `/api/listCarousel`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "GetCategory",
			Router: `/api/listCategories`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "GetCategory2",
			Router: `/api/listGiftCategories`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "ApiCode",
			Router: `/api/code`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "UserLogin",
			Router: `/username/login`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "UserLoginAction",
			Router: `/username/login/action`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "ListProduct",
			Router: `/api/listProduct`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "ThisPruoduct",
			Router: `/api/product`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "HotProducts",
			Router: `/api/hotProducts`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "ListMarket",
			Router: `/api/product/listProduct/:id`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "ShoppingCart",
			Router: `/shopping-cart`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "Message",
			Router: `/message`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "Center",
			Router: `/center`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "Appointment",
			Router: `/appointment`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "Product",
			Router: `/product`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "ProductInfo",
			Router: `/product_info`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "IsLogin",
			Router: `/isLogin`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "CMenu",
			Router: `/createwxmenu`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "Connect",
			Router: `/weixin_connect`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "UserAddress",
			Router: `/api/userAddress`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/web:WebController"] = append(beego.GlobalControllerRouter["ebook/controllers/web:WebController"],
		beego.ControllerComments{
			Method: "UserCar",
			Router: `/api/userCar`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

}
