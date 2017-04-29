package admin

import (
	"ebook/models"
	"github.com/astaxie/beego"
	"time"
)

//@router /api/createMessage [*]
func (this *AjaxController) CreateMessage(){

	userinfo:=this.GetSession("userinfo")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	content:=this.GetString("content")
	user:=userinfo.(models.User)
	var message models.Message
	message.User=&user
	message.Content=content
	message.CreateTime=time.Now()
	if _,err:=message.Insert();err!=nil{
		beego.Error("message Insert Error:",err)
		this.ReturnJson(10001,"留言失败，请重试！")
	}else{
		this.ReturnSuccess()
	}
}

//@router /message/readIt [post]
func (this *AjaxController) ReadIt(){

	mid,_:=this.GetInt64("mid")
	message:=models.Message{Id:mid}
	message.Read()
	message.IsRead = 1
	if _,err:=message.Update();err!=nil{
		beego.Error("message read Error:",err)
		this.ReturnJson(10001,"操作失败")
	}else{
		this.ReturnSuccess()
	}
}
