package controllers

import (
	"github.com/astaxie/beego"
	"ebook/models"
)


type AngularController struct{
	beego.Controller
	IsLogin bool
	User models.User
}

// @router /* [*]
func (this *AngularController) AngularRender(){


	if userinfo:=this.GetSession("users");userinfo!=nil{
		this.TplName="angular/index.html"
		return


	}else{
		this.Ctx.Redirect(302,"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxfcb057b3c57cee69&redirect_uri=http://ebook.hnhqjk.com/api/user/loginAuto&response_type=code&scope=snsapi_userinfo&state=1#wechat_redirect")
		return
	}

}