package models

import "github.com/astaxie/beego/orm"

/*
Status = 0 普通产品分类
Status = 1 特殊产品分类
*/
type Category struct {
	Id int64
	Name string `orm:"unique"`
	Status int    `orm:"default(0)"`
	Image string
	Products []*Product `orm:"reverse(many)"`
}


func (this *Category) Insert ()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *Category) Read()(error){
	return orm.NewOrm().QueryTable("category").Filter("id",this.Id).One(this)
}

func (this *Category) Delete ()(int64,error){
	return orm.NewOrm().Delete(this)
}


func (this *Category) Update ()(int64,error){
	return orm.NewOrm().Update(this)
}
