package web

import (
	"ebook/models"
	"github.com/astaxie/beego/orm"
)

//@router /api/userCar [*]
func (this *WebController) UserCar(){


	userinfo:=this.GetSession("userinfo")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
	}
	user:=userinfo.(models.User)
	var product_info []*models.ProductInfo
	orm.NewOrm().QueryTable("product_info").Filter("user_id",user.Id).All(&product_info)
	this.ReturnSuccess("car_products",product_info)




}