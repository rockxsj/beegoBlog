package controllers

import (
	"beegoBlog/models"
	"beegoBlog/utils"
	//"github.com/astaxie/beego/toolbox"
)

var ModelOptions models.Options

func (this *AdminController) GetOptions() {
	this.Data["Page_title"] = "博客设置"
	this.Data["Is_option"] = true
	this.Layout = "admin/layout.tpl"
	this.TplNames = "admin/option.tpl"
}

func (this *AdminController) SetOptions() {
	inputs := this.Input()
	inputsKey := make([]string, 0, 1000)
	for option, _ := range inputs {
		inputsKey = append(inputsKey, option)
	}
	tableKey := ModelOptions.GetOptionsKey()

	canKey := utils.SliceIntersect(inputsKey, tableKey)

	modOptions := make(map[string]string)
	for _, option := range canKey {
		modOptions[option] = inputs[option][0]
	}

	ModelOptions.SetOptions(modOptions)
	this.Redirect("/admin/option", 302)
}
