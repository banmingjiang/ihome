package controllers

import (
	"github.com/astaxie/beego"
)

type HouseController struct {
	beego.Controller
}

// func (a *HouseController) RetData(resp map[string]interface{}) {
// 	a.Data["json"] = resp
// 	a.ServeJSON()

// }

func (c *HouseController) GetHouseindex() {
	c.Ctx.WriteString("来了老弟")
}
