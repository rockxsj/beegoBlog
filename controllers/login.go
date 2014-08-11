package controllers

import (
	"beegoBlog/models"
	"beegoBlog/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/astaxie/beego/validation"
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
	toolbox.Display("u", username)
	toolbox.Display("pa", password)

	password = utils.LoginPassword(password)
	toolbox.Display("paaa", password)

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
	this.Data["Options"] = ModelOptions.GetOptions()
	this.TplNames = "login/login.tpl"
}
