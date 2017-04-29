package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Collection struct {
	Id int64
	ProductId int64    `orm:"null"`
	ProductHref string `orm:"null"`
	User *User 	   `orm:"rel(fk)"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
}


func (this *Collection) Insert ()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *Collection) Read()(error){
	return orm.NewOrm().QueryTable("collection").Filter("id",this.Id).One(this)
}

func (this *Collection) Delete ()(int64,error){
	return orm.NewOrm().Delete(this)
}


func (this *Collection) Update ()(int64,error){
	return orm.NewOrm().Update(this)
}
