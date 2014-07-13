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
	this.Data["Options"] = ModelOptions.GetOptions()
	this.Data["Cates"] = modelCates.GetCates()
	this.Data["MonthList"] = posts.GetMonthList()
	this.Data["Links"] = ModelLinks.GetLinks()
}

/**
 * 获取所有记录
 */
func (this *PostController) Get() {
	page, _ := this.GetInt(":page")

	cnt := posts.Count()

	var getStart int64
	this.Data["NoPre"], this.Data["NoNext"], this.Data["PrePage"], this.Data["NextPage"], getStart = utils.PageInfo(page, models.INDEX_PRE_NUM, cnt)
	this.Data["Is_index"] = true
	this.Data["BaseUrl"] = "/list"
	ret := posts.GetAll(getStart, models.INDEX_PRE_NUM)
	this.Data["Posts"] = ret
	this.TplNames = "index.tpl"
}

/**
 * 读取一个博文
 *
 * @废弃
 */
func (this *PostController) ShowOne() {
	id, _ := this.GetInt(":id")
	ret := posts.GetOne(id)
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

/**
 * 根据分类获取列表
 */
func (this *PostController) ListByCate() {
	cid, _ := this.GetInt(":cid")
	page, _ := this.GetInt(":page")
	var getStart int64
	cnt := posts.CateCount(cid)
	this.Data["NoPre"], this.Data["NoNext"], this.Data["PrePage"], this.Data["NextPage"], getStart = utils.PageInfo(page, models.INDEX_PRE_NUM, cnt)
	this.Data["Posts"] = posts.GetByCate(cid, getStart, models.INDEX_PRE_NUM)
	this.Data["Cid"] = cid
	this.Data["BaseUrl"] = "/cate/" + strconv.Itoa(int(cid))
	this.TplNames = "index.tpl"
}

/**
 * 根据月份获取列表
 */
func (this *PostController) ListByMonth() {
	month := this.GetString(":month")
	page, _ := this.GetInt(":page")
	var getStart int64
	cnt := posts.MonthCount(month)
	this.Data["NoPre"], this.Data["NoNext"], this.Data["PrePage"], this.Data["NextPage"], getStart = utils.PageInfo(page, models.INDEX_PRE_NUM, cnt)
	this.Data["Posts"] = posts.GetByMonth(month, getStart, models.INDEX_PRE_NUM)
	this.Data["BaseUrl"] = "/month/" + month
	this.TplNames = "index.tpl"
}
