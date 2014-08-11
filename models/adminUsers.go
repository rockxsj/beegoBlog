package models

import (
	"beegoBlog/utils"
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
)

/**
 * 更改密码
 */
func (this *Users) ChangePassword(username, oldPassword, password string) error {
	if !this.CheckLogin(username, utils.LoginPassword(oldPassword)) {
		return errors.New("check old password fail")
	}
	password = utils.LoginPassword(password)
	toolbox.Display("u", username)
	toolbox.Display("n", password)
	num, err := o.QueryTable(this).Filter("username", username).Update(orm.Params{"password": password})
	toolbox.Display("err", err)
	toolbox.Display("num", num)
	return err
}
