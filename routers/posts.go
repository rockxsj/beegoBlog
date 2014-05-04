package routers

import (
	"github.com/astaxie/beego"
	"rockxsj/controllers"
)

func init() {
	beego.Router("/", &controllers.PostController{})
	beego.Router("/post/:id:int", &controllers.PostController{}, "get:ShowOne")
}
