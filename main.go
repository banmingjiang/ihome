package main

import (
	_ "ihome/routers"
	"net/http"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
)

func main() {
	//ignoreStaticPath()

	beego.Run()

}

func ignoreStaticPath() {
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url:", orpath)
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
	//http.ServeFile
}
