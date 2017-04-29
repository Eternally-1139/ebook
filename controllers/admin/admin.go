package admin

import (
	"github.com/astaxie/beego/orm"
	"ebook/models"
)

type AdminController struct {
	Base
}

// @router /admin/home
func (this *AdminController) Home(){
	this.TplName="admin/index.tpl"
}

// @router /admin/category
func (this *AdminController) Category(){

	var category []*models.Category
	orm.NewOrm().QueryTable("category").All(&category)
	this.Data["category"] = category
	this.TplName="admin/category.tpl"
}


// @router /admin/product
func (this *AdminController) Product(){
	var category []*models.Category
	orm.NewOrm().QueryTable("category").All(&category)
	this.Data["category"] = category
	var product []*models.Product
	orm.NewOrm().QueryTable("product").RelatedSel().All(&product)
	this.Data["product"] = product
	this.TplName="admin/product.tpl"
}

// @router /admin/order_finished
func (this *AdminController) Finished(){

	var order_finished []*models.Order
	orm.NewOrm().QueryTable("order").Filter("status",2).RelatedSel().All(&order_finished)
	this.Data["order_finished"] = order_finished

	this.TplName="admin/order_finished.tpl"
}

// @router /admin/order_unfinished
func (this *AdminController) Unfinished(){

	var order_unfinished []*models.Order
	orm.NewOrm().QueryTable("order").Filter("status",1).RelatedSel().All(&order_unfinished)

	this.Data["order_unfinished"] = order_unfinished
	this.TplName="admin/order_unfinished.tpl"
}

// @router /admin/message
func (this *AdminController) Message(){
	var message []*models.Message

	orm.NewOrm().QueryTable("message").RelatedSel().OrderBy("-id").All(&message)
	this.Data["message"] = message
	this.TplName="admin/message.tpl"
}

// @router /admin/comment
func (this *AdminController) Comment(){
	this.TplName="admin/comment.tpl"
}

// @router /admin/appointment
func (this *AdminController) Appointment(){

	var appointment []*models.Appointment

	orm.NewOrm().QueryTable("appointment").RelatedSel().OrderBy("-id").All(&appointment)
	this.Data["appointment"] = appointment
	this.TplName="admin/appointment.tpl"
}

// @router /admin/carousel
func (this *AdminController) Carousel(){

	var carousel []*models.Carouse
	orm.NewOrm().QueryTable("carouse").All(&carousel)
	var products []*models.Product
	orm.NewOrm().QueryTable("product").RelatedSel().All(&products)
	this.Data["products"] = products
	this.Data["carousel"] = carousel
	this.TplName="admin/carousel.tpl"
}