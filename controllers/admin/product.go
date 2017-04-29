package admin

import (
	"ebook/models"
	"github.com/astaxie/beego"
	"time"
	"github.com/astaxie/beego/orm"
	"ebook/service"
)


// @router /admin/deleteProduct [post]
func (this *AjaxController) DeleteProduct(){

	proId ,_:= this.GetInt64("id")
	product := models.Product{Id:proId}

	if _,err:=product.Delete();err!=nil{
		this.ReturnJson(10001,"product insert error")
	}else{
		this.ReturnSuccess()
	}

}

// @router /admin/product/upload [post]
func (this *AjaxController) ProductUpload(){


	no:=string(service.Krand(16,1)) //生成随机商品编号
	product:=models.Product{No:no}
	if err:=product.FindByNo();err==nil{
		this.ReturnJson(10002,"后台生成商品编号重复，请重新添加该产品！")
	}
	product.No=no
	product.Content=this.GetString("proContent")
	product.Name=this.GetString("proName")
	product.Image=this.GetString("cover")
	proPrice ,_:= this.GetFloat("proPrice")
	id,_:=this.GetInt64("proType")
	product.Price=proPrice
	product.CreateTime=time.Now()
	category:=models.Category{Id:id}
	if err:=category.Read();err==nil{
		product.Category=&category
	}


	if _,err:=product.Insert();err!=nil{
		beego.Error("Product Insert Error:",err)
		this.ReturnJson(10001,"product insert error")
	}else{
		this.ReturnSuccess()
	}

}

// @router /admin/picUpload [post]
func (this *AjaxController) UploadPic(){

	var pic models.Pic
	pic.Image=this.GetString("imgSrc")
	id,_:=this.GetInt64("proId")


	product:=models.Product{Id:id}
	if err:=product.Read();err==nil{
		pic.Product=&product
	}
	if _,err:=pic.Insert();err!=nil{
		beego.Error("pic Insert Error:",err)
		this.ReturnJson(10001,"pic insert error")
	}else{
		this.ReturnSuccess()
	}

}

// @router /admin/closeUpload [post]
func (this *AjaxController) CloseUpload(){

	id,_:=this.GetInt64("proId")
	product:=models.Product{Id:id}
	if err:=product.Read();err!=nil{
		this.ReturnJson(10002,"product Read Error")
		return
	}
	product.IsImage = 1
	if _,err:=product.Update();err!=nil{
		beego.Error("product Update Error:",err)
		this.ReturnJson(10001,"product update error")
	}else{
		this.ReturnSuccess()
	}

}

// @router /admin/changeProduct [post]
func (this *AjaxController) ChangeProduct(){

	proId ,_:= this.GetInt64("changeId")
	product := models.Product{Id:proId}
	if err:=product.Read();err!=nil{
		this.ReturnJson(10002,"product Read Error")
		return
	}
	product.Name = this.GetString("changeName")
	changePrice ,_:= this.GetFloat("changePrice")
	product.Price = changePrice
	product.Content = this.GetString("changeContent")
	product.Image = this.GetString("cover2")
	if _,err:=product.Update();err!=nil{
		beego.Error("product Update Error:",err)
		this.ReturnJson(10001,"product update error")
	}else{
		this.ReturnSuccess()
	}

}


// @router /admin/product/showProduct [*]
func (this *AjaxController) ShowProduct(){
	id,_:=this.GetInt64("proId")
	result := make(map[string]interface{})
	var pic []*models.Pic
	orm.NewOrm().QueryTable("pic").Filter("Product",id).RelatedSel().All(&pic)
	result["status"]=10000
	result["pic"]=pic
	this.Data["json"] = result
	this.ServeJSON()

}

