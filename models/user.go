package models

import (
    "github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id            int64
	Mobile        string           `orm:"null"`
	Name          string           `orm:"null"`
	Username      string 	       `orm:"null"`
	Password      string           `orm:"null"`
	Account       float32          `orm:"default(0)"`
	Mark	      int64            `orm:"default(0)"`
	HeadImage     string	       `orm:"null"`
	Sex 	      string	       `orm:"null"`
	OpenId        string           `orm:"null"`
	LastLoginTime time.Time        `orm:"type(datetime)"`
	LastLoginIp   string
	CreateTime    time.Time        `orm:"type(datetime);auto_now_add"`
	CreateIp      string
	Messages      []*Message       `orm:"reverse(many);null"`
	Collection    []*Collection    `orm:"reverse(many);null"`
	Appointments  []*Appointment   `orm:"reverse(many);null"`
	Address       []*Address       `orm:"reverse(many);null"`
	SelectAddress string	       `orm:"null"`
	ProductInfo   []*ProductInfo   `orm:"reverse(many);null"`
	Order         []*Order	       `orm:"reverse(many);null"`
}

func (this *User) FindByMobile() error {
	return orm.NewOrm().QueryTable("user").Filter("mobile", this.Mobile).RelatedSel().One(this)
}

func (this *User) FindByUserName() error {
	return orm.NewOrm().QueryTable("user").Filter("username", this.Username).RelatedSel().One(this)
}

func (this *User) FindByOpenId() error {
	return orm.NewOrm().QueryTable("user").Filter("open_id", this.OpenId).RelatedSel().One(this)
}


func (this *User) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *User) Read() error {
	return orm.NewOrm().QueryTable("user").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *User) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

