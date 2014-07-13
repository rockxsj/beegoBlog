package controllers

import (
	"beegoBlog/models"
	//"github.com/astaxie/beego/toolbox"
)

var modelCates models.Cates

/**
 * 获取分类
 */
func (this *AdminController) GetCates() {
	this.Data["Cates"] = modelCates.GetCates()
	this.Data["Is_cate"] = true
	this.TplNames = "admin/cate.tpl"
	this.Layout = "admin/layout.tpl"
}

/**
 * 添加分类
 */
func (this *AdminController) AddCate() {
	name := this.GetString("cate")
	modelCates.AddCate(name)
	this.Redirect("/admin/GetCates", 302)
}

/**
 * 删除分类
 */
func (this *AdminController) DelCate() {
	ret := make(map[string]int)
	id, _ := this.GetInt("id")
	err := modelCates.DelCate(id)
	if err == nil {
		ret["success"] = 1
	} else {
		ret["success"] = 0
	}
	this.Data["json"] = ret
	this.ServeJson()
	this.EnableRender = false
}

/**
 * 更改分类名
 */
func (this *AdminController) RenameCate() {
	id, _ := this.GetInt("id")
	newName := this.GetString("NewName")
	err := modelCates.RenameCate(id, newName)
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
