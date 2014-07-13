package routers

import (
	"beegoBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin", &controllers.AdminController{}, "get:DashBoard")

	beego.Router("/admin/list", &controllers.AdminController{}, "get:List")
	beego.Router("/admin/list/:page:int", &controllers.AdminController{}, "get:List")

	beego.Router("/admin/add", &controllers.AdminController{}, "get:Add")
	beego.Router("/admin/do_add", &controllers.AdminController{}, "post:DoAdd")

	beego.Router("/admin/update/:id:int", &controllers.AdminController{}, "get:Update")
	beego.Router("/admin/do_update/:id:int", &controllers.AdminController{}, "post:DoUpdate")

	beego.Router("/admin/del/:id:int", &controllers.AdminController{}, "get:Del")

	beego.Router("/admin/option", &controllers.AdminController{}, "get:GetOptions")
	beego.Router("/admin/SetOptions", &controllers.AdminController{}, "post:SetOptions")

	beego.Router("/admin/GetCates", &controllers.AdminController{}, "get:GetCates")
	beego.Router("/admin/AddCate", &controllers.AdminController{}, "post:AddCate")
	beego.Router("/admin/DelCate", &controllers.AdminController{}, "post:DelCate")
	beego.Router("/admin/RenameCate", &controllers.AdminController{}, "post:RenameCate")

	beego.Router("/admin/GetLinks", &controllers.AdminController{}, "get:GetLinks")
	beego.Router("/admin/AddLink", &controllers.AdminController{}, "post:AddLink")
	beego.Router("/admin/DelLink", &controllers.AdminController{}, "post:DelLink")

	beego.Router("/admin/ShowUsers", &controllers.AdminController{}, "get:ShowUsers")
	beego.Router("/admin/ChangePassword", &controllers.AdminController{}, "post:ChangePassword")
}
