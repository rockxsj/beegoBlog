package controllers

/**
 * 用户管理
 */
func (this *AdminController) ShowUsers() {
	this.Data["Is_users"] = true
	this.Layout = "admin/layout.tpl"
	this.TplNames = "admin/users.tpl"
}

/**
 * 更改用户密码
 */
func (this *AdminController) ChangePassword() {
	username := this.GetSession("username").(string)
	oldPassword := this.GetString("oldPassword")
	password := this.GetString("newPassword")
	this.EnableRender = false
	ret := make(map[string]int)
	if user.ChangePassword(username, oldPassword, password) == nil {
		ret["success"] = 1
	} else {
		ret["success"] = 0
	}
	this.Data["json"] = ret
	this.ServeJson()
}
