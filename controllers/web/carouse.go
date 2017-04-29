package web

import (
	"ebook/models"
	"github.com/astaxie/beego/orm"
)

//@router /api/listCarousel [*]
func (this *WebController) ListCarousel(){

	var carouse []*models.Carouse
	orm.NewOrm().QueryTable("carouse").All(&carouse)

	outputs:=make([]map[string]interface{},0)

	for _,swipe:=range carouse{
		output:=make(map[string]interface{})
		output["id"]=swipe.Id
		output["title"]=swipe.Title
		output["img"]=swipe.Img
		output["url"]=swipe.Url
		outputs=append(outputs,output)
	}

	this.ReturnSuccess("swipe",outputs)
}
