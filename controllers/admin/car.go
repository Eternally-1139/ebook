package admin

import (
	"ebook/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//@router /api/addCar [*]
func (this *AjaxController) AddCar(){
	
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
