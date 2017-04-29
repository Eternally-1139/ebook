package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id int64
	Product *Product `orm:"rel(fk);null"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	CommentUserName string
	CommentUserAvatar string
	Content string `orm:"default(0)"`


}


func (this *Comment) Insert()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *Comment) Update()(int64,error){
	return orm.NewOrm().Update(this)
}

func (this *Comment) Delete()(int64,error){
	return orm.NewOrm().Delete(this)
}

func (this *Comment)Read()error{
	return orm.NewOrm().QueryTable("comment").Filter("id",this.Id).RelatedSel().One(this)
}