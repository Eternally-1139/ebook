package admin

import (
	"ebook/models"
	"github.com/astaxie/beego/orm"
)

//@router /admin/addCarousel [*]
func (this *AjaxController) AddCarousel(){



	content:=this.GetString("content")
	image:=this.GetString("image")
	href:=this.GetString("href")

	var carouse models.Carouse
	cnt ,_:= orm.NewOrm().QueryTable("carouse").Count()

	if cnt>=5{
		this.ReturnJson(10002,"最大只能添加五张轮播图！")
		return
	}
	carouse.Title = content
	carouse.Img = image
	carouse.Url = href

	if _,err:=carouse.Insert();err!=nil{
		this.ReturnJson(10001,"插入轮播图失败")
		return
	}else{
		this.ReturnSuccess()
	}


}

//@router /admin/deleteCarousel [*]
func (this *AjaxController) DeleteCarousel(){

	id ,_:=this.GetInt64("id")
	carouse := models.Carouse{Id:id}
	carouse.Delete()
	this.ReturnSuccess()

}