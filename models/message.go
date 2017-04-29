package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

/*
IsRead 0未读 1已读
*/

type Message struct {
	Id int64
	Content string `orm:"null"`
	IsRead int `orm:"default(0)"`
	User *User `orm:"rel(fk)"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
}


func (this *Message) Insert ()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *Message) Read()(error){
	return orm.NewOrm().QueryTable("message").Filter("id",this.Id).One(this)
}

func (this *Message) Delete ()(int64,error){
	return orm.NewOrm().Delete(this)
}


func (this *Message) Update ()(int64,error){
	return orm.NewOrm().Update(this)
}
