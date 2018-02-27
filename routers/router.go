package routers

import (
	"code.aliyun.com/goddream/will/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test", &controllers.MainController{}, "get:Test")
    beego.Router("/login", &controllers.MainController{}, "get:Login")
}
