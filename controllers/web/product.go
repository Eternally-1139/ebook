package web

import (
	"github.com/astaxie/beego/orm"
	"ebook/models"
	"strconv"
	"github.com/astaxie/beego/utils/pagination"
)

//@router /api/listProduct [*]
func (this *WebController) ListProduct(){

	var products []*models.Product
	cid,_:=this.GetInt64("cid")
	orm.NewOrm().QueryTable("product").Filter("category_id",cid).RelatedSel().All(&products)
	this.ReturnSuccess("products",products)
}

//@router /api/product [*]
func (this *WebController) ThisPruoduct(){
	pid,_:=this.GetInt64("pid")
	var pic []*models.Pic
	orm.NewOrm().QueryTable("pic").Filter("product_id",pid).All(&pic)
	var product models.Product
	orm.NewOrm().QueryTable("product").Filter("id",pid).One(&product)

	outputs := make([]map[string]interface{},0)
	output:=make(map[string]interface{})
	output["Id"]=product.Id
	output["Price"]=product.Price
	output["CreateTime"]=product.CreateTime
	output["Name"]=product.Name
	output["Image"]=product.Image
	output["Buy"]=product.Buy
	output["Content"]=product.Content
	output["Pic"]=pic

	outputs=append(outputs,output)
	this.ReturnSuccess("product",output)
}

//@router /api/hotProducts [*]
func (this *WebController) HotProducts(){

	var products []*models.Product
	orm.NewOrm().QueryTable("product").OrderBy("-Buy").Limit(10).RelatedSel().All(&products)
	this.ReturnSuccess("products",products)

}

// @router /api/product/listProduct/:id [*]
func (this *WebController) ListMarket(){
	pers := 6
	category, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"),10,64)
	qs := orm.NewOrm().QueryTable("product")
	if category != 0 {
		qs = orm.NewOrm().QueryTable("product").Filter("category_id", category)
	}
	cnt, _ := models.CountObjects(qs)
	pager := pagination.NewPaginator(this.Ctx.Request, pers, cnt)
	qs = qs.OrderBy("-CreateTime").Limit(pers, pager.Offset()).RelatedSel()
	var products []models.Product
	models.ListObjects(qs, &products)

	outputs := make([]map[string]interface{},0)
	for _, product := range products {
		output:=make(map[string]interface{})
		output["Id"]=product.Id
		output["Price"]=product.Price
		output["CreateTime"]=product.CreateTime
		output["Name"]=product.Name
		output["Image"]=product.Image
		output["Buy"]=product.Buy
		output["Content"]=product.Content

		outputs=append(outputs,output)
	}

	this.ReturnSuccess("products",outputs,"page",pager.Page(),"hasNext",pager.HasNext(),"cnt",cnt,"pers",pers)
}