package controllers

import (
	"github.com/astaxie/beego"

)


type AngularController struct{
	beego.Controller
}

// @router /* [*]
func (this *AngularController) VueRender(){


	if userinfo:=this.GetSession("userinfo");userinfo!=nil{

		this.Ctx.Redirect(302,"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxfcb057b3c57cee69&redirect_uri=http://ebook.hnhqjk.com/api/user/loginAuto&response_type=code&scope=snsapi_userinfo&state=1#wechat_redirect")
		return 

	}else{
		this.TplName="angular/index.html"
	}
}