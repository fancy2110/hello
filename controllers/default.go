package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Test() {
	v := c.GetSession("asta")
	echo_item := "hello"
    if v == nil {
		c.SetSession("asta", int(1))
		v = 0
		echo_item = "zero"
    } else {
        c.SetSession("asta", v.(int)+1)
	}

	c.Ctx.WriteString(echo_item + " | "+ fmt.Sprintf("%d", v))
}


func (c *MainController) Login() {
	c.Ctx.WriteString("Login Page!");
}
