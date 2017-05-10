package admin

import (
	"ebook/models"
	"github.com/astaxie/beego"
)

// @router /admin/addCategory
func (this *AjaxController) AddCategory(){

	cateName := this.GetString("cateName")
	cateType ,_:= this.GetInt("cateType")
	image := this.GetString("image")
	category := models.Category{}

	category.Name = cateName
	category.Status = cateType
	category.Image = image
	if _,err:=category.Insert();err!=nil{
		this.ReturnJson(10001,"category insert error")
	}else{
		this.ReturnSuccess()
	}

}
// @router /admin/deleteCategory
func (this *AjaxController) DeleteCategory(){

	cateId ,_:= this.GetInt64("id")
	category := models.Category{Id:cateId}

	if _,err:=category.Delete();err!=nil{
		this.ReturnJson(10001,"category insert error")
	}else{
		this.ReturnSuccess()
	}

}
// @router /admin/changeCategory [post]
func (this *AjaxController) ChangeCategory(){
	cateId ,_:= this.GetInt64("id")
	category := models.Category{Id:cateId}
	if err:=category.Read();err!=nil{
		this.ReturnJson(10002,"category Read Error")
		return
	}
	category.Name = this.GetString("changeName")

	if _,err:=category.Update();err!=nil{
		beego.Error("category Update Error:",err)
		this.ReturnJson(10001,"category update error")
	}else{
		this.ReturnSuccess()
	}

}



