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

//获取用户信息
func (c *LoginController) GetUserInfo() {
	//先判断用户是否已经登录
	isLogin := c.GetSession("IsLogin")

	if isLogin == true {
		c.Data["IsLogin"] = isLogin
		c.Data["real_name"] = c.GetSession("real_name")
		c.Data["mobile"] = c.GetSession("mobile")
		c.Data["id"] = c.GetSession("id")
	}

}

//渲染登录页面
func (c *LoginController) Login() {
	//进入该页面先判断用户是否已经登录
	getLogin := c.GetSession("IsLogin")
	if getLogin == true {
		c.Ctx.Redirect(303, "/")
	}

	c.TplName = "login.html"
}

//登录方法
func (c *LoginController) ToLog() {
	resp := make(map[string]interface{})
	data := make(map[string]interface{})
	defer c.RetData(resp)
	defer c.RetData(data)
	//获取前端post json数据
	json.Unmarshal(c.Ctx.Input.RequestBody, &resp)

	//查询数据库是否存在对应数据
	o := orm.NewOrm()
	user := models.User{}
	real_name := resp["rel_name"]
	beego.Info("real_name", real_name)
	err := o.Raw("SELECT * FROM user WHERE real_name = ? or mobile=?", real_name, real_name).QueryRow(&user)
	if err != nil {

		data["errno"] = "500"
		data["errmsg"] = "账号不存在"
		models.MakeLogs(data["errmsg"].(string), err)
		return
	}

	beego.Info("user info is:", user)

	//验证哈希加密
	passwd := GetSHA256HashCode([]byte(resp["password"].(string)))

	if passwd != user.Password_hash {
		data["errno"] = "500"
		data["errmsg"] = "密码错误"
		models.MakeLogs(data["errmsg"].(string), "500")
		return
	}
	//验证通过存session
	c.SetSession("real_name", user.Real_name)
	c.SetSession("mobile", user.Mobile)
	c.SetSession("id", user.Id)
	c.SetSession("IsLogin", true)
	data["errno"] = "200"
	data["errmsg"] = "登陆成功"
	models.MakeLogs("用户登录成功", user.Real_name+user.Mobile)

}
func (c *LoginController) LogOut() {
	resp := make(map[string]interface{})

	defer c.RetData(resp)
	c.DelSession("real_name")
	c.DelSession("mobile")
	c.DelSession("id")
	c.SetSession("IsLogin", false)
	resp["errno"] = "0"
	resp["errmsg"] = "退出成功"

}
