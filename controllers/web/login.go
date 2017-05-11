package web

import (
	"ebook/models"
	"github.com/astaxie/beego/orm"
	"github.com/dchest/captcha"
)

// @router /api/code [*]
func (this *WebController) ApiCode(){
	//从验证码服务获取验证码ID
	captchaId := captcha.NewLen(5)
	//生成json数据
	result := make(map[string]interface{})
	//状态
	result["status"] = 10000
	//验证码图片地址
	result["src"] = "/captcha/" + captchaId + ".png"
	result["id"] = captchaId
	this.Data["json"] = result
	this.ServeJSON()
	return
}

// @router /username/login [*]
func (this *WebController) UserLogin() {

	this.TplName="login.html"
}

// @router /username/login/action [*]
func (this *WebController)  UserLoginAction(){

	username:=this.GetString("username")
	password:=this.GetString("password")

	var user models.User
	if err:=orm.NewOrm().QueryTable("user").
		Filter("username",username).Filter("password",password).One(&user);err!=nil{
		this.ReturnJson(10001,"Username or Password Not Found")
		return
	}

	this.ReturnSuccess()

}
