package routers

import (
	"beegoBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PostController{})
	beego.Router("/post/:id:int", &controllers.PostController{}, "get:ShowOne")
	beego.Router("/list/:page:int", &controllers.PostController{}, "get:Get")
}
