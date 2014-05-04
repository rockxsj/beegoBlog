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
func GetAll() []*Posts {
	o := orm.NewOrm()

	var posts []*Posts
	o.QueryTable("posts").All(&posts)

	return posts
}
