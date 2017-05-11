package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Home",
			Router: `/admin/home`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Category",
			Router: `/admin/category`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Product",
			Router: `/admin/product`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Finished",
			Router: `/admin/order_finished`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Unfinished",
			Router: `/admin/order_unfinished`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Message",
			Router: `/admin/message`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Comment",
			Router: `/admin/comment`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Appointment",
			Router: `/admin/appointment`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Carousel",
			Router: `/admin/carousel`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "AddAddress",
			Router: `/api/addAddress`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteAddress",
			Router: `/api/deleteAddress`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "ChangeAddress",
			Router: `/api/changeAddress`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "SelectAddress",
			Router: `/api/selectAddress`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/file/upload`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CreateOrder",
			Router: `/order/createOrder`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "FinishingOrder",
			Router: `/order/finishingOrder`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "ShowUnfinishedOrder",
			Router: `/order/showUnfinishedOrder`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteOrder",
			Router: `/order/deleteOrder`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "GetCar",
			Router: `/api/getCar`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "AddCar",
			Router: `/api/addCar`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteCar",
			Router: `/api/deleteCar`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "AddCarousel",
			Router: `/admin/addCarousel`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteCarousel",
			Router: `/admin/deleteCarousel`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "AddCategory",
			Router: `/admin/addCategory`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteCategory",
			Router: `/admin/deleteCategory`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "ChangeCategory",
			Router: `/admin/changeCategory`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CreateMessage",
			Router: `/api/createMessage`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "ReadIt",
			Router: `/message/readIt`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteProduct",
			Router: `/admin/deleteProduct`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "ProductUpload",
			Router: `/admin/product/upload`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "UploadPic",
			Router: `/admin/picUpload`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CloseUpload",
			Router: `/admin/closeUpload`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "ChangeProduct",
			Router: `/admin/changeProduct`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "ShowProduct",
			Router: `/admin/product/showProduct`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/admin`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "AdminLogin",
			Router: `/admin/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["ebook/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "AdminLogout",
			Router: `/admin/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
