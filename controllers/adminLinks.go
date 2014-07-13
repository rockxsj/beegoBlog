package controllers

import (
	"beegoBlog/models"
	//"beegoBlog/utils"
	//"github.com/astaxie/beego/toolbox"
)

var ModelLinks models.Links

/**
 * 获取友情链接
 */
func (this *AdminController) GetLinks() {
	this.Data["Is_link"] = true
	this.Data["Links"] = ModelLinks.GetLinks()
	this.Layout = "admin/layout.tpl"
	this.TplNames = "admin/link.tpl"
}

/**
 * 增加友情链接
 */
func (this *AdminController) AddLink() {
	name := this.GetString("name")
	link := this.GetString("link")
	ModelLinks.AddLink(name, link)
	this.Redirect("/admin/GetLinks", 302)
}

/**
 * 删除友情链接
 */
func (this *AdminController) DelLink() {
	id, _ := this.GetInt("id")
	err := ModelLinks.DelLink(id)
	ret := make(map[string]int)
	if err == nil {
		ret["success"] = 1
	} else {
		ret["success"] = 0
	}
	this.Data["json"] = ret
	this.ServeJson()
	this.EnableRender = false
}
