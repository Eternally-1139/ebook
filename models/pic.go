package models

import (
	"github.com/astaxie/beego/orm"
)

type Pic struct {
	Id int64
	Product *Product `orm:"rel(fk);null"`
	Image string
}

func (this *Pic) Insert()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *Pic) Update()(int64,error){
	return orm.NewOrm().Update(this)
}

func (this *Pic) Delete()(int64,error){
	return orm.NewOrm().Delete(this)
}

func (this *Pic)Read()error{
	return orm.NewOrm().QueryTable("pic").Filter("id",this.Id).RelatedSel().One(this)
}