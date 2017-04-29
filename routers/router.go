package routers

import (
	"ebook/controllers"
	"github.com/astaxie/beego"
	"ebook/controllers/admin"
	"ebook/controllers/web"
)

func init() {
	beego.SetStaticPath("/web-static", "web-static")
	beego.SetStaticPath("/assets", "assets")
	beego.Include(
		&admin.AdminController{},&controllers.ApiController{},&admin.AjaxController{},&web.WebController{},&controllers.AngularController{})
}
