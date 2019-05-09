package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"ihome/models"
	"text/template"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (a *RegisterController) RetData(resp map[string]interface{}) {
	a.Data["json"] = resp
	a.ServeJSON()
}
func (c *RegisterController) Register() {
	//进入该页面先判断用户是否已经登录
	getLogin := c.GetSession("IsLogin")
	if getLogin == true {
		c.Ctx.Redirect(303, "/")
	}
	c.TplName = "reg.html"
}

func (this *RegisterController) Reg() {

	resp := make(map[string]interface{})
	data := make(map[string]interface{})
	//defer this.RetData(resp)
	defer this.RetData(data)
	//获取前端post json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	// fmt.Println(resp)
	// beego.Info("real_name", resp["rel_name"])
	// beego.Info("password", resp["password"])
	// beego.Info("mobile", resp["mobile"])
	//c.Ctx.WriteString("good")

	o := orm.NewOrm()
	user := models.User{}

	user.Real_name = template.HTMLEscapeString(resp["rel_name"].(string))
	user.Mobile = template.HTMLEscapeString(resp["mobile"].(string))
	//先查看是否已存在该用户名或手机号码
	err := o.Raw("SELECT * FROM user WHERE real_name=? or real_name=?", user.Real_name, user.Mobile).QueryRow(&user)

	if fmt.Sprintf("%s", err) != "<QuerySeter> no row found" {
		data["errno"] = "505"
		data["errmsg"] = "用户名或手机号码已存在"
		models.MakeLogs(data["errmsg"].(string), err)
		return
	}

	//哈希加密接收的是数组类型 先把字符串转成数组类型
	user.Password_hash = GetSHA256HashCode([]byte(resp["password"].(string)))
	id, err := o.Insert(&user)
	if err != nil {
		data["errno"] = "500"
		data["errmsg"] = "网络错误，请重试"
		beego.Info("用户注册失败，err=", err)
		return
	}

	data["errno"] = "200"
	data["errmsg"] = "注册成功"
	models.MakeLogs("用户注册成功，ID=", id)
	//注册成功 存入session
	this.SetSession("real_name", user.Real_name)
	this.SetSession("mobile", user.Mobile)
	this.SetSession("id", id)
	this.SetSession("IsLogin", true)

}

//SHA256生成哈希值
func GetSHA256HashCode(message []byte) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode

	//方法二：
	//bytes2:=sha256.Sum256(message)//计算哈希值，返回一个长度为32的数组
	//hashcode2:=hex.EncodeToString(bytes2[:])//将数组转换成切片，转换成16进制，返回字符串
	//return hashcode2
}
