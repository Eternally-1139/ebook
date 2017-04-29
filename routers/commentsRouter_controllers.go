package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["shopktl/controllers:AngularController"] = append(beego.GlobalControllerRouter["shopktl/controllers:AngularController"],
		beego.ControllerComments{
			Method: "VueRender",
			Router: `/*`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["shopktl/controllers:ApiController"] = append(beego.GlobalControllerRouter["shopktl/controllers:ApiController"],
		beego.ControllerComments{
			Method: "UserAutoLogin",
			Router: `/api/user/autoLogin`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["shopktl/controllers:ApiController"] = append(beego.GlobalControllerRouter["shopktl/controllers:ApiController"],
		beego.ControllerComments{
			Method: "UserRegister",
			Router: `/api/userRegister`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["shopktl/controllers:ApiController"] = append(beego.GlobalControllerRouter["shopktl/controllers:ApiController"],
		beego.ControllerComments{
			Method: "GetUserInfo",
			Router: `/api/getUserInfo`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["shopktl/controllers:ApiController"] = append(beego.GlobalControllerRouter["shopktl/controllers:ApiController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/test/getUser`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["shopktl/controllers:ApiController"] = append(beego.GlobalControllerRouter["shopktl/controllers:ApiController"],
		beego.ControllerComments{
			Method: "GetProduct",
			Router: `/api/getProduct`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["shopktl/controllers:ApiController"] = append(beego.GlobalControllerRouter["shopktl/controllers:ApiController"],
		beego.ControllerComments{
			Method: "GetNumbers",
			Router: `/api/getProductAndCategory`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

}
