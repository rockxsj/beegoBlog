package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Prepare() {
	username := this.GetSession("username")

	if username == nil {
		this.Redirect("/login", 302)
	}
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

/**
 * 文章列表
 */
func (this *AdminController) List() {
	this.Data["Page_title"] = "文章列表"
	this.Data["Is_list"] = true
	ret := posts.GetAll()
	for _, v := range ret {
		v.Content = beego.Substr(beego.Html2str(v.Content), 0, 200) + "..."
	}
	this.Data["Posts"] = ret
	this.Layout = "admin/layout.tpl"
	this.TplNames = "admin/list.tpl"
}

/**
 * 打开编辑器，新增文章
 */
func (this *AdminController) Add() {
	this.Data["Is_add"] = true
	this.Data["Do_action"] = "do_add"
	this.Data["Page_title"] = "新增文章"
	this.Layout = "admin/layout.tpl"
	this.TplNames = "admin/edit.tpl"
}

/**
 * 新增文章
 */
func (this *AdminController) DoAdd() {
	title := this.GetString("title")
	content := this.GetString("content")

	ret := posts.AddOne(title, content)
	if ret == true {
		this.Redirect("/admin/list", 302)
	} else {
		this.Ctx.WriteString("insert err")
	}
}

/**
 * 删除文章
 */
func (this *AdminController) Del() {
	id := this.GetString(":id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		this.Ctx.WriteString("id wrong")
	}
	posts.Del(intid)
	this.Redirect("/admin", 302)
}

/**
 * 获取一个，供更新使用
 */
func (this *AdminController) Update() {
	id := this.GetString(":id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		this.Ctx.WriteString("id wrong")
		this.StopRun()
	}
	this.Data["Is_list"] = true
	this.Data["Do_action"] = "do_update/" + id

	ret := posts.GetOne(intid)
	this.Data["Post"] = ret

	this.Data["Page_title"] = ret.Title + "——更新文章"
	this.Layout = "admin/layout.tpl"
	this.TplNames = "admin/edit.tpl"
}

/**
 * 更新
 */
func (this *AdminController) DoUpdate() {
	id := this.GetString(":id")

	intid, err := strconv.Atoi(id)

	if err != nil {
		this.Ctx.WriteString("id wrong")
	}

	title := this.GetString("title")
	content := this.GetString("content")

	posts.DoUpdateOne(intid, title, content)
	this.Redirect("/admin/list", 302)
}
