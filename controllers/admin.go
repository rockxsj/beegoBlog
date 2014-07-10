package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Prepare() {
	username := this.GetSession("username")

	if username == nil {
		this.Redirect("/login", 302)
	}
	this.Data["Options"] = ModelOptions.GetOptions()
}

/**
 * 管理面板
 */
func (this *AdminController) DashBoard() {
	this.Data["Page_title"] = "博客后台"
	this.Data["Is_index"] = true
	this.Layout = "admin/layout.tpl"
	this.TplNames = "admin/index.tpl"
}
