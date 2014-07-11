package controllers

import (
	"beegoBlog/models"
	"beegoBlog/utils"
	"github.com/astaxie/beego"
	"strconv"
)

/**
 * 文章列表
 */
func (this *AdminController) List() {
	this.Data["Page_title"] = "文章列表"
	this.Data["Is_list"] = true
	page, _ := strconv.Atoi(this.GetString(":page"))
	cnt := posts.Count()

	var getStart int
	this.Data["NoPre"], this.Data["NoNext"], this.Data["PrePage"], this.Data["NextPage"], getStart = utils.PageInfo(page, models.ADMIN_PRE_NUM, int(cnt))
	ret := posts.GetAll(getStart, models.ADMIN_PRE_NUM)
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
	this.Data["Cates"] = modelCates.GetCates()
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
	cid, _ := this.GetInt("cid")

	ret := posts.AddOne(title, content, cid)
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
	this.Data["Cates"] = modelCates.GetCates()
	id := this.GetString(":id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		this.Ctx.WriteString("id wrong")
		this.StopRun()
	}
	this.Data["Is_update"] = true
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
	cid, _ := this.GetInt("cid")

	posts.DoUpdateOne(intid, title, content, cid)
	this.Redirect("/admin/list", 302)
}
