package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"rockxsj/models"
)

type LoginController struct {
	beego.Controller
}

var user models.Users

/**
 * 登录
 */
func (this *LoginController) Login() {
	val := validation.Validation{}
	username := this.GetString("username")
	password := this.GetString("password")

	val.Required(username, "username")
	val.Required(password, "password")

	if val.HasErrors() {
		for _, v := range val.Errors {
			this.Ctx.WriteString("<script>alert('" + v.Key + " " + v.Message + "');</script>") //遇到一个错误就可以StopRun了其实，这里只是练习验证模块
		}
		this.StopRun()
	}

	md5_hash := md5.New()
	md5_hash.Write([]byte(password))
	password = hex.EncodeToString(md5_hash.Sum(nil))

	if user.CheckLogin(username, password) == true {
		this.SetSession("username", username)
		this.Ctx.SetCookie("username", username)
		this.Redirect("/admin", 302)
	} else {
		this.Redirect("/login", 302)
	}
}

/**
 * 登出
 */
func (this *LoginController) Logout() {
	this.DestroySession()
	this.Redirect("/", 302)
}

/**
 * 登陆面板
 */
func (this *LoginController) ShowLoginBox() {
	if this.GetSession("username") != nil {
		this.Redirect("/admin", 302)
		this.StopRun()
	}
	this.TplNames = "login/login.tpl"
}
