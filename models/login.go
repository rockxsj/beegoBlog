package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Users struct {
	Id       int
	Username string
	Password string
	Email    string
	Nickname string
	Avatar   string
	Addtime  time.Time
	ModTime  time.Time
}

func init() {
	orm.RegisterModel(new(Users))
	orm.RegisterDataBase("user", "mysql", beego.AppConfig.String("mysql_conf"))
	orm.Debug = true
}

/**
 * 检查登陆
 */
func CheckLogin(username string, password string) bool {
	o := orm.NewOrm()

	var user Users
	err := o.Raw("SELECT * FROM users WHERE username = ? AND password = ?", username, password).QueryRow(&user)

	if err == nil {
		return true
	} else {
		return false
	}
}
