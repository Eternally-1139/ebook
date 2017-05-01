package web

import (
	"github.com/astaxie/beego/orm"
	"ebook/models"
	"fmt"
	"sort"
	"crypto/sha1"
	"io"
	"strings"
	"ebook/service"
)

type WebController struct {
	Base
}

const Token = "yifangege"

/*
func (this *WebController) Home(){
	var carouse []*models.Carouse
	orm.NewOrm().QueryTable("carouse").All(&carouse)

	var category1 []*models.Category
	orm.NewOrm().QueryTable("category").Filter("status",0).Limit(7).RelatedSel().All(&category1)

	var category2 []*models.Category
	orm.NewOrm().QueryTable("category").Filter("status",1).Limit(6).RelatedSel().All(&category2)

	this.Data["category1"]=category1
	this.Data["category2"]=category2
	this.Data["carouse"] = carouse
	this.TplName="index.tpl"
}
*/
//@router /shopping-cart
func (this *WebController) ShoppingCart(){
	this.TplName="cart.tpl"
}

//@router /message
func (this *WebController) Message(){
	this.TplName="message.tpl"
}

//@router /center
func (this *WebController) Center(){
	this.TplName="user.tpl"
}

//@router /appointment
func (this *WebController) Appointment(){
	this.TplName="appointment.tpl"
}

//@router /product
func (this *WebController) Product(){
	var products []*models.Product
	cid,_:=this.GetInt64("cid")
	orm.NewOrm().QueryTable("product").Filter("category_id",cid).RelatedSel().All(&products)
	this.Data["products"] = products
	this.TplName="product.tpl"
}

//@router /product_info
func (this *WebController) ProductInfo(){
	var pic []*models.Pic
	pid,_:=this.GetInt64("pid")
	orm.NewOrm().QueryTable("pic").Filter("product_id",pid).RelatedSel().All(&pic)
	var product models.Product
	orm.NewOrm().QueryTable("product").Filter("id",pid).RelatedSel().One(&product)
	this.Data["pic"] = pic
	this.Data["product"] = product
	this.TplName="product_info.tpl"
}

//@router /isLogin [*]
func (this *WebController) IsLogin(){

	userinfo:=this.GetSession("userinfo")
	if userinfo==nil{
		this.ReturnJson(10403,"请先登录")
		return
	}else {
		this.ReturnSuccess()
	}


}

//@router /createwxmenu [*]
func (this *WebController) CMenu(){
	ac:=service.GetAccessToken("wxfcb057b3c57cee69","218f0ea06e24651010db6a1f0eb8f40c");
	service.CreateWxMenu(ac)
	this.ReturnSuccess()

}

//@router /weixin_connect [*]
func (this *WebController) Connect(){


	//微信接入验证 这是首次对接微信 填写url后 微信服务器会发一个请求过来
	//c.Ctx.Request.URL-------------wx_connect?signature=038d75ed5485b9881a01b3b93e85f9fff28ea739&echostr=5756456183388806654&timestamp=1476173150&nonce=1093541731

	//开发者提交信息(包括URL、Token)后，微信服务器将发送Http Get请求到填写的URL上，
	//GET请求携带四个参数：signature、timestamp、nonce和echostr。公众号服务程序应该按如下要求进行接入验证
	timestamp, nonce,signatureIn := this.GetString("timestamp"), this.GetString("nonce"),this.GetString("signature")
	signatureGen := makeSignature(timestamp, nonce)

	//将加密后获得的字符串与signature对比，如果一致，说明该请求来源于微信
	if signatureGen != signatureIn {
		fmt.Printf("signatureGen != signatureIn signatureGen=%s,signatureIn=%s\n", signatureGen, signatureIn)
		this.Ctx.WriteString("")

	} else {
		//如果请求来自于微信，则原样返回echostr参数内容 以上完成后，接入验证就会生效，开发者配置提交就会成功。
		echostr := this.GetString("echostr")
		this.Ctx.WriteString(echostr)
	}

	this.TplName="index.tpl"
}

func makeSignature(timestamp, nonce string) string {

	//1. 将 plat_token、timestamp、nonce三个参数进行字典序排序
	sl := []string{Token, timestamp, nonce}
	sort.Strings(sl)
	//2. 将三个参数字符串拼接成一个字符串进行sha1加密
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))

	return fmt.Sprintf("%x", s.Sum(nil))
}
