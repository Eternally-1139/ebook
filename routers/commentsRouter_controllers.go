package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ebook/controllers:AngularController"] = append(beego.GlobalControllerRouter["ebook/controllers:AngularController"],
		beego.ControllerComments{
			Method: "AngularRender",
			Router: `/*`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers:ApiController"] = append(beego.GlobalControllerRouter["ebook/controllers:ApiController"],
		beego.ControllerComments{
			Method: "LoginDo",
			Router: `/api/user/loginAuto`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers:ApiController"] = append(beego.GlobalControllerRouter["ebook/controllers:ApiController"],
		beego.ControllerComments{
			Method: "UserAutoLogin",
			Router: `/api/user/autoLogin`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers:ApiController"] = append(beego.GlobalControllerRouter["ebook/controllers:ApiController"],
		beego.ControllerComments{
			Method: "GetUserInfo",
			Router: `/api/getUserInfo`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers:ApiController"] = append(beego.GlobalControllerRouter["ebook/controllers:ApiController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/test/getUser`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers:ApiController"] = append(beego.GlobalControllerRouter["ebook/controllers:ApiController"],
		beego.ControllerComments{
			Method: "GetProduct",
			Router: `/api/getProduct`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers:ApiController"] = append(beego.GlobalControllerRouter["ebook/controllers:ApiController"],
		beego.ControllerComments{
			Method: "GetNumbers",
			Router: `/api/getProductAndCategory`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

}
