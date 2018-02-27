package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

/**
* 用户模块，负责用户的注册，登录以及基础信息管理
*/
type UserController struct {
	beego.Controller
}

// Login  登录接口
// @param username :string, 用户名
// @param password :string, 用户密码
// @param vcode	  :string, 验证码， 非必须
// @return 
func (c *UserController) Login() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// Register 注册接口, 完成用户的创建以及基础信息的配置
func (c *UserController) Register() {
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

// Forget 忘记密码
func (c *UserController) Forget() {
	c.Ctx.WriteString("Login Page!");
}

// Profile 个人信息展示页面
func (c *UserController) Profile() {

}
