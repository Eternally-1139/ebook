package models

import "github.com/astaxie/beego/orm"

type Carouse struct {
	Id int64
	Title string `orm:"null"`
	Img string   `orm:"null"`
	Url  string   `orm:"null"`
}


func (this *Carouse) Insert ()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *Carouse) Read()(error){
	return orm.NewOrm().QueryTable("carouse").Filter("id",this.Id).One(this)
}

func (this *Carouse) Delete ()(int64,error){
	return orm.NewOrm().Delete(this)
}


func (this *Carouse) Update ()(int64,error){
	return orm.NewOrm().Update(this)
}
