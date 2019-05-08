package controllers

import (
	"encoding/json"
	"ihome/models"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

type AreaController struct {
	beego.Controller
}

func (a *AreaController) RetData(resp map[string]interface{}) {
	a.Data["json"] = resp
	a.ServeJSON()
}

func (c *AreaController) GetArea() {
	resp := make(map[string]interface{})
	//捕捉异常
	defer c.RetData(resp)
	//c.Ctx.WriteString("i'm api")
	//从数据库获取数据
	//获取数据结构
	area := []models.Area{}
	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&area)
	if err != nil || num == 0 {
		resp["errno"] = "4000"
		resp["errmsg"] = "获取数据失败"
		return
	}
	resp["errno"] = "0"
	resp["errmsg"] = "OK"
	resp["data"] = &area

	json_data, err := json.Marshal(resp)
	if err != nil {
		resp["errno"] = 40002
		resp["errmsg"] = "change json data err"

		return
	}
	//c.RetData(resp)
	beego.Info(json_data)

}
