package admin

import (
	"github.com/astaxie/beego"
	"ebook/models"
)

type Base struct {
	beego.Controller
	IsLogin bool
	User models.User
}

type CtrPrepare interface {
	CtrPrepare()
}

func (this *Base) Prepare(){

	this.IsLogin=false

	if userinfo:=this.GetSession("userinfo");userinfo!=nil{

		user:=userinfo.(models.User)
		if err:=user.Read();err!=nil{
			this.Ctx.Redirect(302,"/")
			return
		}else{
			if user.Id==1{
				this.User=user
				this.IsLogin=true
				this.Data["User"]=user
			}else{
				this.Ctx.Redirect(302,"/admin")
				return
			}

		}

	}

	if !this.IsLogin{
		this.Ctx.Redirect(302,"/admin")
		return
	}

	if app, ok := this.AppController.(CtrPrepare); ok {
		app.CtrPrepare()
	}



}