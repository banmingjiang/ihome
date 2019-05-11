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
	//SetFDFS()
	//models.Upload("main.go")
	beego.Run()

}

//设置fdfs的静态文件路径
func SetFDFS() {
	beego.SetStaticPath("group1/MOO", "fdfs/storange_data/data")
}

//自定义视图层
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

//自定义错误页面
