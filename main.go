package main

import (
	_ "code.aliyun.com/goddream/will/routers"
	_ "github.com/astaxie/beego/session/mysql"
	_ "code.aliyun.com/goddream/will/filters"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

