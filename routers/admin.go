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
}
