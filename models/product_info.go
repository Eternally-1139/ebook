package models

import (
	"github.com/astaxie/beego/orm"
)

type ProductInfo struct {
	Id int64
	User *User `orm:"rel(fk);null"`
	Image string
	Name string
	Content string
	Num  int   `orm:"default(1)"`
	Price float64
	ProductId int64 `orm:"default(0)"`
}

func (this *ProductInfo) Insert()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *ProductInfo) Update()(int64,error){
	return orm.NewOrm().Update(this)
}

func (this *ProductInfo) Delete()(int64,error){
	return orm.NewOrm().Delete(this)
}

func (this *ProductInfo)Read()error{
	return orm.NewOrm().QueryTable("product_info").Filter("id",this.Id).RelatedSel().One(this)
}
func (this *ProductInfo) FindById() error {
	return orm.NewOrm().QueryTable("product_info").Filter("product_id", this.ProductId).RelatedSel().One(this)
}
