package admin

import (
	"ebook/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//@router /api/getCar [*]
func (this *AjaxController) GetCar(){
	userinfo:=this.GetSession("users")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	user:=userinfo.(models.User)
	user.Read()
	var cart *[]models.ProductInfo
	orm.NewOrm().QueryTable("product_info").Filter("user_id",user.Id).One(&cart)
	this.ReturnSuccess("cart",cart)
}

//@router /api/addCar [*]
func (this *AjaxController) AddCar(){
	userinfo:=this.GetSession("users")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	user:=userinfo.(models.User)
	user.Read()
	id,_:=this.GetInt64("id")
	name:=this.GetString("name")
	price,_:=this.GetFloat("price")
	img:=this.GetString("img")
	content:=this.GetString("content")

	productInfo:=models.ProductInfo{ProductId:id}
	if err:=productInfo.FindById();err==nil{
		productInfo.Read()
		productInfo.Num+=1
		productInfo.Update()
		this.ReturnSuccess()
		return
	}else{
		productInfo.Name=name
		productInfo.Price=price
		productInfo.Image=img
		productInfo.Num=1
		productInfo.Content=content
		productInfo.Insert()
		this.ReturnSuccess()
		return
	}
}

//@router /api/deleteCar [*]
func (this *AjaxController) DeleteCar(){

	userinfo:=this.GetSession("users")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	product_no:=this.GetString("product_no")
	product:=models.Product{No:product_no}
	product.Read()

	no:=product.No

	user:=userinfo.(models.User)
	var product_info models.ProductInfo
	user.Read()

	orm.NewOrm().QueryTable("product_info").Filter("no",no).One(&product_info)

	if _,err:=product_info.Delete();err!=nil{
		beego.Error("product_info delete Error:",err)
		this.ReturnJson(10001,"从购物车删除失败！")
	}else{
		this.ReturnSuccess()
	}

}
