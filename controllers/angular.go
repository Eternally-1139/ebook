package controllers

import (
	"github.com/astaxie/beego"
	"ebook/models"
	"fmt"
)


type AngularController struct{
	beego.Controller
	IsLogin bool
	User models.User
}

// @router /* [*]
func (this *AngularController) AngularRender(){


	if userinfo:=this.GetSession("users");userinfo!=nil{
		user:=userinfo.(models.User)
		if err:=user.Read();err!=nil{
			this.SetSession("users",user)
			this.IsLogin = true
		}

	}

	if !this.IsLogin{
		this.Ctx.Redirect(302,"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxfcb057b3c57cee69&redirect_uri=http://ebook.hnhqjk.com/api/user/loginAuto&response_type=code&scope=snsapi_userinfo&state=1#wechat_redirect")
		return
	}else{
		user:=this.GetSession("users").(models.User)
		fmt.Println("已登录用户："+user.Name)
		this.TplName="angular/index.html"
	}


}