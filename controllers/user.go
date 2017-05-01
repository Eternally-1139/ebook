package controllers

import (
	"ebook/models"
	"ebook/service"
	"github.com/astaxie/beego"
	"time"
	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

type CodeToken struct {
	Openid string `json:"openid"`
	AccessToken string `json:"access_token"`
}
//@router /api/user/loginAuto [*]
func (this *ApiController) LoginDo(){


	code:=this.GetString("code")
	fmt.Println("code:"+code)
	requestLine:="https://api.weixin.qq.com/sns/oauth2/access_token?appid=wxfcb057b3c57cee69&secret=218f0ea06e24651010db6a1f0eb8f40c&code="+code+"&grant_type=authorization_code"
	resp, err := http.Get(requestLine)

	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("发送get请求获取 openid 错误", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("发送get请求获取 atoken 读取返回body错误", err)
	}

	if bytes.Contains(body, []byte("access_token")) {
		ctk := CodeToken{}
		err = json.Unmarshal(body, &ctk)
		if err != nil {
			fmt.Println("发送get请求获取 atoken 返回数据json解析错误", err)
		}

		fmt.Println("access_token",ctk.AccessToken)
		fmt.Println("openid",ctk.Openid)
		Openid:=ctk.Openid
		user:=models.User{OpenId:Openid}
		userinfo:=service.GetUserInfo(ctk.AccessToken,Openid)
		if(userinfo.Nickname==""){
			fmt.Println("用户信息为空")
			this.ReturnJson(10002,"自动注册失败！请重试")
			return
		}
		if err:=user.FindByOpenId();err==nil{
			user.Read()
			this.SetSession("users",user)
			this.Ctx.Redirect(302,"/")
			return
		}else{

			user.Name=userinfo.Nickname
			user.Sex=userinfo.Sex
			user.HeadImage=userinfo.Headimgurl
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
				this.Ctx.Redirect(302,"/")
				return
			}
		}

		//createWxMenu(atr.AccessToken)
		defer resp.Body.Close()

	} else {

	}
	this.ReturnJson(10003,"access_token获取失败！")

}
//@router /api/user/autoLogin [*]
func (this *ApiController) UserAutoLogin(){

	requestLine:="https://open.weixin.qq.com/connect/qrconnect?appid="+"wxfcb057b3c57cee69"+"&redirect_uri="+"http://ebook.hnhqjk.com/api/user/loginAuto"+"&response_type=code&scope="+"snsapi_userinfo"+"&state=STATE#wechat_redirect"
	resp, err := http.Get(requestLine)

	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("发送get请求获取 openid 错误", err)
	}else{
		fmt.Println("发送请求成功！")
		this.Ctx.Redirect(302,"/api/user/loginAuto")
	}
	this.ReturnSuccess()

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