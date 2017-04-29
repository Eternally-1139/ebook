package admin

import (
	"ebook/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//@router /api/addCar [*]
func (this *AjaxController) AddCar(){

	userinfo:=this.GetSession("userinfo")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	productId,_:=this.GetInt64("productId")
	product:=models.Product{Id:productId}
	product.Read()
	user:=userinfo.(models.User)
	product_info2 := models.ProductInfo{No:product.No}
	var product_info models.ProductInfo
	if err:=product_info2.FindByNo();err==nil{
		product_info2.Num=product_info2.Num+1
		product_info2.Update()
		beego.Debug("product_num:",product_info2.Num)

		this.ReturnSuccess()
		return
	}
	product_info.Name = product.Name
	product_info.Price = product.Price
	product_info.Image = product.Image
	product_info.Num=1
	product_info.No = product.No
	product_info.User=&user



	if _,err:=product_info.Insert();err!=nil{
		beego.Error("product_info Insert Error:",err)
		this.ReturnJson(10001,"添加购物车失败！")
	}else{
		this.ReturnSuccess()
	}

}

//@router /api/deleteCar [*]
func (this *AjaxController) DeleteCar(){

	userinfo:=this.GetSession("userinfo")
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
