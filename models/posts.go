package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	INDEX_PRE_NUM = 3  //前台每页文章数量
	ADMIN_PRE_NUM = 10 //后台每页文章数量
)

type Posts struct {
	Id       int64
	Title    string
	Content  string
	Add_time time.Time
	Mod_time time.Time
	Cid      int64
}

/**
 * 获取所有记录
 */
func (this *Posts) GetAll(start, limit int64) []*Posts {
	var posts []*Posts
	o.QueryTable("posts").Limit(limit, start).OrderBy("-id").All(&posts)

	return posts
}

/**
 * 获取一个记录
 */
func (this *Posts) GetOne(id int64) *Posts {
	this.Id = id
	o.Read(this)
	return this
}

/**
 * 获取相邻记录
 *
 * @param int 当前id
 * @param bool 之前还是之后
 */
func (this *Posts) GetPreOrNext(id int64, next bool) (error, *Posts) {
	var fiter, order string
	if next {
		fiter = "Id__gt"
		order = "Id"
	} else {
		fiter = "Id__lt"
		order = "-Id"
	}
	preOrNext := new(Posts)
	err := o.QueryTable(this).Filter(fiter, id).OrderBy(order).Limit(1).One(preOrNext)
	if err == orm.ErrNoRows {
		return orm.ErrNoRows, preOrNext
	} else {
		return nil, preOrNext
	}
}

/**
 * 获取当前表中文章数
 */
func (this *Posts) Count() (cnt int64) {
	cnt, err := o.QueryTable(this).Count()
	if err != nil {
		panic("查询表出错")
	}
	return
}

/**
 * 根据分类获取记录
 */
func (this *Posts) GetByCate(cate, start, limit int64) []*Posts {
	var posts []*Posts
	o.QueryTable("posts").Limit(limit, start).Filter("Cid", cate).OrderBy("-id").All(&posts)

	return posts
}

/**
 * 根据分类查询数量
 */
func (this *Posts) CateCount(cate int64) (cnt int64) {
	cnt, err := o.QueryTable(this).Filter("Cid", cate).Count()
	if err != nil {
		panic("查询表出错")
	}
	return
}

/**
 * 取出月份
 */
func (this *Posts) GetMonthList() orm.ParamsList {
	var monthList orm.ParamsList
	o.Raw("select distinct(date_format(add_time, '%Y-%m')) as each_month from posts").ValuesFlat(&monthList)
	toolbox.Display("o", monthList)
	return monthList
}

/**
 * 根据月份获取文章数
 */
func (this *Posts) MonthCount(month string) (cnt int64) {
	cnt, err := o.QueryTable(this).Filter("add_time__startswith", month).Count()
	if err != nil {
		panic("查询表出错")
	}
	return
}

/**
 * 根据月份获取文章
 */
func (this *Posts) GetByMonth(month string, start, limit int64) []*Posts {
	var posts []*Posts
	o.QueryTable("posts").Limit(limit, start).Filter("add_time__startswith", month).OrderBy("-id").All(&posts)
	return posts
}
