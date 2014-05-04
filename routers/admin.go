package routers

import (
	"github.com/astaxie/beego"
	"rockxsj/controllers"
)

func init() {
	beego.Router("/admin", &controllers.AdminController{}, "get:DashBoard")

	beego.Router("/admin/list", &controllers.AdminController{}, "get:List")

	beego.Router("/admin/add", &controllers.AdminController{}, "get:Add")
	beego.Router("/admin/do_add", &controllers.AdminController{}, "post:DoAdd")

	beego.Router("/admin/update/:id:int", &controllers.AdminController{}, "get:Update")
	beego.Router("/admin/do_update/:id:int", &controllers.AdminController{}, "post:DoUpdate")

	beego.Router("/admin/del/:id:int", &controllers.AdminController{}, "get:Del")
}
