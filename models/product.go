package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


/*
	Image 封面，Pic 产品详情图
	IsImage 0未上传产品详情图 1已上传产品详情图
*/
type Product struct {
	Id int64
	Price float64 `orm:"default(0)"`
	Category *Category `orm:"rel(fk);null"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	Name string
	No   string   `orm:"unique"`
	Buy  int      `orm:"default(0)"`
	Image string
	IsImage int  	`orm:"default(0)"`
	Pic []*Pic `orm:"reverse(many)"`
	Comment []*Comment `orm:"reverse(many)"`
	Content string

}

func (this *Product) FindByNo() error {
	return orm.NewOrm().QueryTable("product").Filter("no", this.No).RelatedSel().One(this)
}


func (this *Product) Insert()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *Product) Update()(int64,error){
	return orm.NewOrm().Update(this)
}

func (this *Product) Delete()(int64,error){
	return orm.NewOrm().Delete(this)
}

func (this *Product)Read()error{
	return orm.NewOrm().QueryTable("product").Filter("id",this.Id).RelatedSel().One(this)
}