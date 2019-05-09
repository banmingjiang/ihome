package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// func (a *HouseController) RetData(resp map[string]interface{}) {
// 	a.Data["json"] = resp
// 	a.ServeJSON()
// }
func (c *MainController) Get() {
	c.Data["real_name"] = nil
	c.Data["IsLogin"] = c.GetSession("IsLogin")
	//判断是否登陆
	//beego.Info("the IsLogin is:", c.Data["IsLogin"])
	if c.Data["IsLogin"] != nil {
		c.Data["IsLogin"] = c.Data["IsLogin"]
		c.Data["real_name"] = c.GetSession("real_name")
		c.Data["mobile"] = c.GetSession("mobile")
	}

	// beego.Info("the session is:", seesione)
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	//c.Layout = "head.html"
	c.TplName = "index.html"
}
