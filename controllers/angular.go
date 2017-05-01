package controllers

import (
	"github.com/astaxie/beego"

	"fmt"
)


type AngularController struct{
	beego.Controller
}

// @router /* [*]
func (this *AngularController) AngularRender(){


	if userinfo:=this.GetSession("users");userinfo!=nil{

		this.Ctx.Redirect(302,"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxfcb057b3c57cee69&redirect_uri=http://ebook.hnhqjk.com/api/user/loginAuto&response_type=code&scope=snsapi_userinfo&state=1#wechat_redirect")
		return

	}else{
		fmt.Println("已登录用户："+this.GetSession("users").(string))
		this.TplName="angular/index.html"
	}
}