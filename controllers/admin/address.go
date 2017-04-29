package admin

import (
	"github.com/astaxie/beego"
	"ebook/models"
)


//@router /api/addAddress [*]
func (this *AjaxController) AddAddress(){

	userinfo:=this.GetSession("userinfo")
	consignee:=this.GetString("consignee")
	mobile:=this.GetString("mobile")
	provence:=this.GetString("provence")
	city:=this.GetString("city")
	county:=this.GetString("county")
	place:=this.GetString("place")
	var address models.Address
	user:=userinfo.(models.User)
	address.User=&user
	address.Province = provence
	address.City = city
	address.County = county
	address.Place = place
	address.Consignee = consignee
	address.Mobile = mobile
	if _,err:=address.Insert();err!=nil{
		beego.Error("address Insert Error:",err)
		this.ReturnJson(10001,"地址添加失败！")
	}else{
		this.ReturnSuccess()
	}


}

//@router /api/deleteAddress [*]
func (this *AjaxController) DeleteAddress(){

	aid,_:=this.GetInt64("aid")
	address:=models.Address{Id:aid}
	if _,err:=address.Delete();err!=nil{
		beego.Error("address Insert Error:",err)
		this.ReturnJson(10001,"收货地址删除失败！")
	}else{
		this.ReturnSuccess()
	}

}

//@router /api/changeAddress [*]
func (this *AjaxController) ChangeAddress(){

	userinfo:=this.GetSession("userinfo")
	aid,_:=this.GetInt64("aid")
	consignee:=this.GetString("consignee")
	mobile:=this.GetString("mobile")
	provence:=this.GetString("provence")
	city:=this.GetString("city")
	county:=this.GetString("county")
	place:=this.GetString("place")
	address:=models.Address{Id:aid}

	user:=userinfo.(models.User)

	address.User=&user

	address.Consignee = consignee
	address.Mobile = mobile
	address.Province = provence
	address.City = city
	address.County = county
	address.Place = place
	if _,err:=address.Update();err!=nil{
		beego.Error("address Update Error:",err)
		this.ReturnJson(10001,"修改收货地址失败！")
	}else{
		this.ReturnSuccess()
	}

}
//@router /api/selectAddress [*]
func (this *AjaxController) SelectAddress(){

	userinfo:=this.GetSession("userinfo")
	aid,_:=this.GetInt64("aid")

	address:=models.Address{Id:aid}
	user:=userinfo.(models.User)
	if err:=address.Read();err!=nil{
		beego.Error("address Read Error:",err)
		this.ReturnJson(10001,"设置默认收货地址失败！")
	}else{
		if err:=user.Read();err!=nil{
			beego.Error("user Read Error:",err)
			this.ReturnJson(10001,"设置默认收货地址失败！")
		}else{
			user.SelectAddress = address.Province+address.County+address.City+address.Place
			user.Update()
			this.ReturnSuccess()
		}
	}

}


