package controllers

import (
	"ebook/models"
	"ebook/service"
	"github.com/astaxie/beego"
	"time"

	"fmt"
)

//@router /api/user/autoLogin [*]
func (this *ApiController) UserAutoLogin(){

	ac:=service.GetAccessToken("wxfcb057b3c57cee69","218f0ea06e24651010db6a1f0eb8f40c")
	Subscribe, Openid, Nickname, Sex, Language, City, Province, Country, Headimgurl:=service.GetUserInfo(ac)
	address:=City+Province+Country
	if(address==""){
		fmt.Println("error")
		return
	}
	beego.Debug("openid"+address)
	user:=models.User{OpenId:Openid}

	if err:=user.FindByOpenId();err==nil{
		user.Read()
		this.SetSession("userinfo",user)
		this.ReturnSuccess()
		return
	}else{

		user.Name=Nickname
		user.Language=Language
		user.Subscribe=Subscribe
		user.Sex=Sex
		user.HeadImage=Headimgurl
		user.CreateIp = "localhost"
		user.Account = 0
		user.CreateTime = time.Now()
		user.LastLoginTime = time.Now()
		user.LastLoginIp = "localhost"
		user.Mark = 0
		user.Mobile = ""
		if _,err:=user.Insert();err!=nil{
			beego.Error("User Insert Error:",err)
			this.ReturnJson(10002,"User Insert Error")
		}else{
			this.SetSession("userinfo",user)
			this.ReturnSuccess()
		}
	}



}

//@router /api/userRegister [*]
func (this *ApiController) UserRegister(){


	var user models.User
	username := this.GetString("username")
	name := this.GetString("nickName")
	sex := this.GetString("sex")
	image := this.GetString("headimgurl")
	password := service.StrToMd5Sha1(this.GetString("password"))

	user.Username=username
	user.Name=name
	user.Sex=sex
	user.HeadImage=image
	user.Password=password
	user.CreateIp = "localhost"
	user.Account = 0
	user.Address = nil
	user.CreateTime = time.Now()
	user.LastLoginTime = time.Now()
	user.LastLoginIp = "localhost"
	user.Mark = 0
	user.Mobile = ""

	if err:=user.FindByUserName();err==nil{
		this.ReturnJson(10001,"Username Has Exist")
		return
	}

	if _,err:=user.Insert();err!=nil{
		beego.Error("User Insert Error:",err)
		this.ReturnJson(10002,"User Insert Error")
	}else{
		this.SetSession("userinfo",user)
		this.ReturnSuccess()
	}
}



//@router /api/getUserInfo [*]
func (this *ApiController) GetUserInfo(){

	userinfo:=this.GetSession("userinfo")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
	}
	user:=userinfo.(models.User)
	result := make(map[string]interface{})
	result["status"]=10000
	result["id"] = user.Id
	result["mobile"] = user.Mobile
	result["name"] = user.Name
	result["account"] = user.Account
	result["mark"] = user.Mark
	result["sex"] = user.Sex
	result["headImage"] = user.HeadImage

	this.Data["json"] = result
	this.ServeJSON()


}