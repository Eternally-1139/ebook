package models

import "github.com/astaxie/beego/orm"

type Address struct {
	Id int64
	Consignee string
	Mobile    string
	Province  string 	`orm:"null"`
	City 	  string	`orm:"null"`
	County    string 	`orm:"null"`
	Place	  string        `orm:"null"`
	User      *User 	`orm:"rel(fk)"`
}


func (this *Address) Insert ()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *Address) Read()(error){
	return orm.NewOrm().QueryTable("address").Filter("id",this.Id).One(this)
}

func (this *Address) Delete ()(int64,error){
	return orm.NewOrm().Delete(this)
}


func (this *Address) Update ()(int64,error){
	return orm.NewOrm().Update(this)
}
