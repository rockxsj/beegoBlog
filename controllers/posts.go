package controllers

import (
	"beegoBlog/models"
	"beegoBlog/utils"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/toolbox"
	"strconv"
)

type PostController struct {
	beego.Controller
}

var posts models.Posts

/**
 * "构造函数"
 */
func (this *PostController) Prepare() {
	username := this.GetSession("username")

	if username == nil {
		this.Data["Is_login"] = false
	} else {
		this.Data["Username"] = username
		this.Data["Is_login"] = true
	}
}

/**
 * 获取所有记录
 */
func (this *PostController) Get() {
	page, _ := strconv.Atoi(this.GetString(":page"))

	cnt := posts.Count()

	var getStart int
	this.Data["NoPre"], this.Data["NoNext"], this.Data["PrePage"], this.Data["NextPage"], getStart = utils.PageInfo(page, models.INDEX_PRE_NUM, int(cnt))
	this.Data["Is_index"] = true
	this.Data["Page_title"] = "首页——博客"
	ret := posts.GetAll(getStart, models.INDEX_PRE_NUM)
	for _, v := range ret {
		v.Content = beego.Substr(beego.Html2str(v.Content), 0, 200) + "..."
	}
	this.Data["Posts"] = ret
	this.Layout = "layout.tpl"
	this.TplNames = "index.tpl"
}

/**
 * 读取一个博文
 */
func (this *PostController) ShowOne() {
	id := this.GetString(":id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		this.Ctx.WriteString("id wrong")
		this.StopRun()
	}
	ret := posts.GetOne(intid)
	this.Data["Page_title"] = ret.Title + "——博客"
	this.Data["Post"] = ret
	err_pre, pre := posts.GetPreOrNext(ret.Id, false)
	err_next, next := posts.GetPreOrNext(ret.Id, true)
	if err_pre == nil {
		this.Data["Pre"] = pre
		this.Data["NoPre"] = false
	} else {
		this.Data["NoPre"] = true
	}
	if err_next == nil {
		this.Data["Next"] = next
		this.Data["NoNext"] = false
	} else {
		this.Data["NoNext"] = true
	}

	this.Layout = "layout.tpl"
	this.TplNames = "post.tpl"
}
