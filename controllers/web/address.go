package web

import (
	"ebook/models"
	"github.com/astaxie/beego/orm"
)

//@router /api/userAddress [*]
func (this *WebController) UserAddress(){

	userinfo:=this.GetSession("userinfo")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
	}
	user:=userinfo.(models.User)

	var address []*models.Address
	orm.NewOrm().QueryTable("address").Filter("user_id",user.Id).All(&address)
	this.ReturnSuccess("address",address)

}
