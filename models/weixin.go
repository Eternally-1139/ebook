package models

import (
	"github.com/astaxie/beego/orm"

	"time"
)

type Weixin struct {

	Id int64
	AccessToken string 	`orm:"null"`
	JsApiTicket string	`orm:"null"`
	GetAcTime time.Time  	`orm:"null"`
	GetTiTime time.Time  	`orm:"null"`
	ExpiresIn   float64  	`orm:"null"`

}


func (this *Weixin) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Weixin) Read() error {
	return orm.NewOrm().QueryTable("weixin").Filter("id", this.Id).RelatedSel().One(this)
}


func (this *Weixin) Update()(int64,error){
	return orm.NewOrm().Update(this)
}

func (this *Weixin) Delete()(int64,error){
	return orm.NewOrm().Delete(this)
}