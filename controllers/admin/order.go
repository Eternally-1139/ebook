package admin

import (
	"ebook/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)


// @router /order/createOrder [*]
func (this *AjaxController) CreateOrder(){

	//需加入购买逻辑

	var price float64
	buyMethod,_:=this.GetInt("buyMethod")
	userinfo:=this.GetSession("userinfo")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}
	payMethod ,_:= this.GetInt("payMethod")
	payType,_:=this.GetInt("payType")
	var order models.Order

	user:=userinfo.(models.User)
	order.User=&user

	if buyMethod==1{

		var product_info []*models.ProductInfo
		orm.NewOrm().QueryTable("product_info").Filter("user_id",user.Id).All(&product_info)


		for j:=0;j<len(product_info);j++{
			//price+=(product_info[j].Price)*(product_info[j].Num)

			if price==0{
				this.ReturnJson(10004,"购物车空空如也")
			}
		}
		fmt.Println(price)
	}else if buyMethod==2{

		pid,_:=this.GetInt64("pid")
		product:=models.Product{Id:pid}
		if err:=product.Read();err!=nil{
			beego.Error("order Insert Error:",err)
			this.ReturnJson(10003,"获取商品信息失败")
		}else{
			price=product.Price
		}

	}else{
		this.ReturnJson(10002,"订单类型错误")
	}

	order.PayMethod = payMethod
	order.Price = price
	order.PayType = payType
	order.Status = 1
	order.CreateIp = "localhost"
	//添加订单收货地址
	order.Address = user.SelectAddress


	if err:=order.Insert();err!=nil{
		beego.Error("order Insert Error:",err)
		this.ReturnJson(10001,"product insert error")
	}else{
		this.ReturnSuccess()
	}


}

// @router /order/finishingOrder [post]
func (this *AjaxController) FinishingOrder(){

	oId ,_:=this.GetInt64("oId")
	oddNumber:=this.GetString("oddNumber")
	order:=models.Order{Id:oId}
	if err:=order.Read();err!=nil{
		this.ReturnJson(10002,"product Read Error")
		return
	}
	order.OddNumber=oddNumber
	order.Status=2
	if _,err:=order.Update();err!=nil{
		beego.Error("order Update Error:",err)
		this.ReturnJson(10001,"order update error")
	}else{
		this.ReturnSuccess()
	}

}
// @router /order/showUnfinishedOrder [post]
func (this *AjaxController) ShowUnfinishedOrder(){

	result := make(map[string]interface{})
	var order_unfinished []*models.Order
	orm.NewOrm().QueryTable("order").Filter("status",1).All(&order_unfinished)
	result["order"] = order_unfinished
	result["status"]=10000
	this.Data["json"] = result
	this.ServeJSON()

}
// @router /order/deleteOrder [post]
func (this *AjaxController) DeleteOrder(){

	oId ,_:= this.GetInt64("oId")
	order := models.Order{Id:oId}

	if _,err:=order.Delete();err!=nil{
		this.ReturnJson(10001,",order delete error")
	}else{
		this.ReturnSuccess()
	}
}


