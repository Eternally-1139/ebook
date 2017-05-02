package models

import (
	"github.com/astaxie/beego/orm"
)

type ProductInfo struct {
	Id int64
	User *User `orm:"rel(fk);null"`
	Image string
	Name string
	Num  float64   `orm:"default(1)"`
	Price float64
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
