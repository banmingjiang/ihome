package controllers

import (
	"encoding/json"
	"fmt"
	"ihome/models"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

type SettingController struct {
	beego.Controller
}

func (a *SettingController) RetData(resp map[string]interface{}) {
	a.Data["json"] = resp
	a.ServeJSON()
}
func (c *SettingController) Setting() {

	c.Data["IsLogin"] = c.GetSession("IsLogin")
	//判断是否登陆
	//beego.Info("the IsLogin is:", c.Data["IsLogin"])
	if c.Data["IsLogin"] != nil {
		c.Data["IsLogin"] = true
		c.Data["id"] = c.GetSession("id")
		data := GetUserInfo(c.Data["id"])
		if data["code"] != "200" {
			c.Ctx.Redirect(303, "/api/login.html")
			return
		}
		res := data["res"].(map[string]interface{})
		c.Data["real_name"] = res["real_name"]
		c.Data["mobile"] = res["mobile"]
		c.Data["avatar_url"] = res["avatar_url"]
		c.Data["name"] = res["name"]

		fmt.Printf("============%v", res)

	} else {
		c.Ctx.Redirect(303, "/api/login.html")
	}

	// beego.Info("the session is:", seesione)
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	//c.Layout = "head.html"
	c.TplName = "setting.html"
}

func (c *SettingController) SetUser() {
	//定义一个map存储转过来的值
	resp := make(map[string]interface{})
	data := make(map[string]interface{})
	defer c.RetData(resp)
	//获取网页端的值存入resp
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {
		resp["code"] = 500
		resp["msg"] = "数据获取失败"
		return
	}

	real_name := data["real_name"]
	user := []models.User{}
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM user WHERE real_name =?", real_name).QueryRows(&user)

	if err != nil {
		resp["code"] = 500
		resp["msg"] = "数据查询失败"

		return
	}
	if num >= 1 {
		resp["code"] = 505
		resp["msg"] = "该昵称已存在"

		return
	} else {
		resp["code"] = 200
		resp["msg"] = "ok"
	}

}
