package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func LoginCheckFilter(ctx * context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
    if !ok && ctx.Request.RequestURI != "/login" {
        ctx.Redirect(302, "/login")
    }
}

func init() {
	beego.InsertFilter("/*", beego.BeforeExec, LoginCheckFilter);
}