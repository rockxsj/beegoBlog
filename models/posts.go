package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	INDEX_PRE_NUM = 3  //前台每页文章数量
	ADMIN_PRE_NUM = 10 //后台每页文章数量
)

type Posts struct {
	Id       int
	Title    string
	Content  string
	Add_time time.Time
	Mod_time time.Time
	Cid      int64
}

/**
 * 获取所有记录
 */
func (this *Posts) GetAll(start, pre int) []*Posts {
	var posts []*Posts
	o.QueryTable("posts").Limit(pre, start).OrderBy("-id").All(&posts)

	return posts
}

/**
 * 获取一个记录
 */
func (this *Posts) GetOne(id int) *Posts {
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
func (this *Posts) GetPreOrNext(id int, next bool) (error, *Posts) {
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
