package controllers

import (
	"github.com/astaxie/beego/orm"
	"ebook/models"
)


type ApiController struct {
	Base
}

//@router /test/getUser [*]
func (this *ApiController) GetUser(){
	result:=make(map[string]interface{})
	user:=models.User{Id:1}
	user.Read()
	result["user"]=user
	this.Data["json"]=result
	this.ServeJSON()
}

//@router /api/getProduct [*]
func (this *ApiController) GetProduct(){

	result := make(map[string]interface{})
	var product []*models.Product
	orm.NewOrm().QueryTable("product").All(&product)
	this.Data["product"] = product
	result["product"]=product
	this.Data["json"] = result
	this.ServeJSON()

}


//@router /api/getProductAndCategory [*]
func (this *ApiController) GetNumbers(){
	result := make(map[string]interface{})
	var product []*models.Product
	orm.NewOrm().QueryTable("product").All(&product)
	result["product"]=product

	var category []*models.Category
	orm.NewOrm().QueryTable("category").All(&category)
	result["category"]=category

	var message []*models.Message
	orm.NewOrm().QueryTable("message").Filter("is_read",0).All(&message)
	result["message"]=message

	var appointment []*models.Appointment
	orm.NewOrm().QueryTable("appointment").Filter("is_read",0).All(&appointment)
	result["appointment"]=appointment

	result["status"]=10000

	this.Data["json"] = result
	this.ServeJSON()
}

