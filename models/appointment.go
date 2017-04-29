package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Appointment struct {
	Id int64
	Name string    `orm:"null"`
	Mobile string  `orm:"null"`
	Time string `orm:"null"`
	Content string `orm:"null"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	User *User `orm:"rel(fk)"`
	IsRead int `orm:"default(0)"`
}


func (this *Appointment) Insert ()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *Appointment) Read()(error){
	return orm.NewOrm().QueryTable("appointment").Filter("id",this.Id).One(this)
}

func (this *Appointment) Delete ()(int64,error){
	return orm.NewOrm().Delete(this)
}


func (this *Appointment) Update ()(int64,error){
	return orm.NewOrm().Update(this)
}
