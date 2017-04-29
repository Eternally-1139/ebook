package web

import (
	"github.com/astaxie/beego/orm"
	"ebook/models"
)

//@router /api/listCategories [*]
func (this *WebController) GetCategory(){

	var category []*models.Category
	orm.NewOrm().QueryTable("category").Filter("status",0).RelatedSel().All(&category)

	this.ReturnSuccess("categories",category)
}

//@router /api/listGiftCategories [*]
func (this *WebController) GetCategory2(){

	var category2 []*models.Category
	orm.NewOrm().QueryTable("category").Filter("status",1).RelatedSel().All(&category2)
	this.ReturnSuccess("categories2",category2)
}
