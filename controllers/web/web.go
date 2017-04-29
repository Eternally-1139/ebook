package web

import (
	"github.com/astaxie/beego/orm"
	"ebook/models"
)

type WebController struct {
	Base
}


/*
func (this *WebController) Home(){
	var carouse []*models.Carouse
	orm.NewOrm().QueryTable("carouse").All(&carouse)

	var category1 []*models.Category
	orm.NewOrm().QueryTable("category").Filter("status",0).Limit(7).RelatedSel().All(&category1)

	var category2 []*models.Category
	orm.NewOrm().QueryTable("category").Filter("status",1).Limit(6).RelatedSel().All(&category2)

	this.Data["category1"]=category1
	this.Data["category2"]=category2
	this.Data["carouse"] = carouse
	this.TplName="index.tpl"
}
*/
//@router /shopping-cart
func (this *WebController) ShoppingCart(){
	this.TplName="cart.tpl"
}

//@router /message
func (this *WebController) Message(){
	this.TplName="message.tpl"
}

//@router /center
func (this *WebController) Center(){
	this.TplName="user.tpl"
}

//@router /appointment
func (this *WebController) Appointment(){
	this.TplName="appointment.tpl"
}

//@router /product
func (this *WebController) Product(){
	var products []*models.Product
	cid,_:=this.GetInt64("cid")
	orm.NewOrm().QueryTable("product").Filter("category_id",cid).RelatedSel().All(&products)
	this.Data["products"] = products
	this.TplName="product.tpl"
}

//@router /product_info
func (this *WebController) ProductInfo(){
	var pic []*models.Pic
	pid,_:=this.GetInt64("pid")
	orm.NewOrm().QueryTable("pic").Filter("product_id",pid).RelatedSel().All(&pic)
	var product models.Product
	orm.NewOrm().QueryTable("product").Filter("id",pid).RelatedSel().One(&product)
	this.Data["pic"] = pic
	this.Data["product"] = product
	this.TplName="product_info.tpl"
}

//@router /isLogin [*]
func (this *WebController) IsLogin(){

	userinfo:=this.GetSession("userinfo")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}else {
		this.ReturnSuccess()
	}

}
