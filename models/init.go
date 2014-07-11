package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	o orm.Ormer
)

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql_conf"))
	orm.RegisterModel(new(Posts))
	orm.RegisterModel(new(Users))
	orm.RegisterModel(new(Options))
	orm.RegisterModel(new(Cates))
	o = orm.NewOrm()
}
