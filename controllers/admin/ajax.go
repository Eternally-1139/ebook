package admin

import (

	"github.com/astaxie/beego"
	"time"
	"ebook/models"
	"ebook/service"
)

type AjaxController struct {
	beego.Controller
}

// @router /admin
func (this *AjaxController) Login(){
	this.TplName="admin/login.tpl"
}

// @router /admin/login [post]
func (this *AjaxController) AdminLogin(){
	username := this.GetString("username")
	password := this.GetString("password")
	user := models.User{Username: username}

	if err := user.FindByUserName(); err != nil {
		this.ReturnJson(10002,"username or password error")
		return
	}

	if user.Password != service.StrToMd5Sha1(password) {
		this.ReturnJson(10002,"username or password error")
		return
	}

	user.LastLoginTime = time.Now()
	user.LastLoginIp = this.Ctx.Input.IP()
	user.Update()
	this.SetSession("userinfo", user)
	this.ReturnSuccess()
}

// @router /admin/logout
func (this *AjaxController) AdminLogout(){
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, "/admin")
}




func (this *AjaxController) ReturnJson(status int, message string, args ...interface{}) {
	result := make(map[string]interface{})
	result["status"] = status
	result["message"] = message

	key := ""

	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}

	this.Data["json"] = result
	this.ServeJSON()

}

func (this *AjaxController) ReturnSuccess(args ...interface{}) {

	result := make(map[string]interface{})
	result["status"] = 10000
	result["message"] = "success"
	key := ""
	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}
	this.Data["json"] = result
	this.ServeJSON()

}