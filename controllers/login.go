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

//从数据库获取用户信息
func GetUserInfo(id interface{}) (data map[string]interface{}) {

	data = make(map[string]interface{})
	res := make(map[string]interface{})
	userinfo := models.User{}
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM user WHERE id = ?", id).QueryRow(&userinfo)
	if err != nil {
		data["code"] = "500"
		data["msg"] = "用户不存在"
		return
	}

	json_data, err := json.Marshal(userinfo)
	if err != nil {
		data["code"] = "505"
		data["msg"] = "json.Marshal ERROR"
		//beego.Info("json.Marshal is fail")
		models.MakeLogs("json.Marshal is fail", err)
		return
	}
	//转成map格式返回
	un_err := json.Unmarshal([]byte(json_data), &res)
	if un_err != nil {
		beego.Info("json.unMarshal is fail")
		models.MakeLogs("json.unMarshal is fail", un_err)
		data["code"] = "505"
		data["msg"] = "json.unMarshal ERROR"
		return
	}
	data["code"] = "200"
	data["msg"] = "OK"
	data["res"] = res
	return
	//h, _ := json.Marshal(data)

	//fmt.Printf("%v\n", string(h))

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
