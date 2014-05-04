package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Posts struct {
	Id       int
	Title    string
	Content  string
	Add_time time.Time
	Mod_time time.Time
}

func init() {
	orm.RegisterModel(new(Posts))
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql_conf"))
}

/**
 * 获取所有记录
 */
func (this *Posts) GetAll() []*Posts {
	o := orm.NewOrm()

	var posts []*Posts
	o.QueryTable("posts").All(&posts)

	return posts
}

/**
 * 获取一个记录
 */
func (this *Posts) GetOne(id int) *Posts {
	o := orm.NewOrm()

	//posts := new(Posts)

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
	o := orm.NewOrm()
	//posts := new(Posts)

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
