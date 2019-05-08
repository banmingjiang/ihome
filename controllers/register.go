package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"ihome/models"

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
	c.TplName = "reg.html"
}

func (this *RegisterController) Reg() {

	resp := make(map[string]interface{})
	defer this.RetData(resp)
	//获取前端post json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	// fmt.Println(resp)
	// beego.Info("real_name", resp["rel_name"])
	// beego.Info("password", resp["password"])
	// beego.Info("mobile", resp["mobile"])
	//c.Ctx.WriteString("good")
	o := orm.NewOrm()
	user := models.User{}
	user.Real_name = resp["rel_name"].(string)
	user.Mobile = resp["mobile"].(string)
	//哈希加密接收的是数组类型 先把字符串转成数组类型
	passwd := []byte(resp["password"].(string))

	user.Password_hash = GetSHA256HashCode(passwd)
	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = models.RECODE_PARAMERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PARAMERR)
		beego.Info("用户注册失败，err=", err)
		return
	}
	beego.Info("用户注册成功，ID=", id)
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	//注册成功 存入session
	this.SetSession("real_name", user.Real_name)
	this.SetSession("mobile", user.Mobile)
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
