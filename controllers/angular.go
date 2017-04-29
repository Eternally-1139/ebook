package controllers

import (
	"github.com/astaxie/beego"
)


type AngularController struct{
	beego.Controller
}

// @router /* [*]
func (this *AngularController) VueRender(){

	this.TplName="angular/index.html"
}