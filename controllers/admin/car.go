package admin

import (
	"ebook/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
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
	var cart []*models.ProductInfo
	orm.NewOrm().QueryTable("product_info").Filter("user_id",user.Id).RelatedSel().All(&cart)
	this.ReturnSuccess("carts",cart)
	return
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
	fmt.Println(user.Name)
	fmt.Println("img"+img)

	var productInfo models.ProductInfo

	orm.NewOrm().QueryTable("product_info").Filter("product_id",id).Filter("user_id",user.Id).One(&productInfo)
	fmt.Println(productInfo)

	if productInfo.ProductId!=0{
		fmt.Println("查询到该用户下该购物车")
		productInfo.Num+=1
		productInfo.Update()
		fmt.Println("商品自增")
		this.ReturnSuccess("info:database add 1","success")
		return
	}else{
		fmt.Println("未查询到该商品")
		var proinfo models.ProductInfo
		proinfo.ProductId=id
		proinfo.Name=name
		proinfo.Price=price
		proinfo.Image=img
		proinfo.Num=1
		proinfo.Content=content
		proinfo.User=&user
		proinfo.Insert()
		this.ReturnSuccess("info:create cart add 1","success")
		return
	}
	fmt.Println("error!")
}


//@router /api/deleteCar [*]
func (this *AjaxController) DeleteCar(){

	userinfo:=this.GetSession("users")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	id,_:=this.GetInt64("id")
	productInfo:=models.ProductInfo{Id:id}
	productInfo.Read()
	if _,err:=productInfo.Delete();err!=nil{
		beego.Error("product_info delete Error:",err)
		this.ReturnJson(10001,"从购物车删除失败！")
	}else{
		this.ReturnSuccess()
	}

}
