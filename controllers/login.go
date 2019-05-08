package controllers

import (
	"encoding/json"
	"ihome/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (a *LoginController) RetData(resp map[string]interface{}) {
	a.Data["json"] = resp
	a.ServeJSON()
}

//渲染登录页面
func (c *LoginController) Login() {
	// seesion := c.GetSession("real_name")
	// beego.Info("the session is:", seesion)
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.html"
}

//登录方法
func (c *LoginController) ToLog() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	//获取前端post json数据
	json.Unmarshal(c.Ctx.Input.RequestBody, &resp)
	beego.Info("real_name", resp["rel_name"])
	beego.Info("password", resp["password"])

	//查询数据库是否存在对应数据
	o := orm.NewOrm()
	user := models.User{}
	user.Real_name = resp["real_name"].(string)
	err := o.Read(&user)
	if err != nil {
		beego.Info("password账号不存在")
		resp["errno"] = "500"
		resp["errmsg"] = "账号不存在"
		return
	}
	beego.Info("找到了数据")
	beego.Info(user)
	//GetSHA256HashCode()
	// seesion := c.GetSession("real_name")
	// beego.Info("the session is:", seesion)
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "login.html"
}
func (c *LoginController) LogOut() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	c.DelSession("real_name")
	c.DelSession("mobile")
	c.SetSession("IsLogin", false)
	resp["errno"] = "0"
	resp["errmsg"] = "退出成功"

}
