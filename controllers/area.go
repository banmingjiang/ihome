package controllers

import (
	"encoding/json"
	"ihome/models"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"

	_ "github.com/astaxie/beego/cache/redis"
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

	defer c.RetData(resp)
	//c.Ctx.WriteString("i'm api")
	//从数据库获取数据
	//获取数据结构
	area := []models.Area{}
	resp["errno"] = "0"
	resp["errmsg"] = "OK"

	//使用菲关系型数据库redis
	redisData, err := cache.NewCache("redis", `{"key":"SetArea","conn":":6379","dbNum":"0"}`)
	if err != nil {
		beego.Info("use redis fail:", err)
		models.MakeLogs("use redis fail:", err)
		return
	}

	//beego.Info("====================================运行到这====================================")
	//如果缓存中有数据，直接返回数据，不再查询数据库
	if redisArea := redisData.Get("AreaJson"); redisArea != nil {
		models.MakeLogs("get redis fail:", redisArea)
		resp["data"] = redisArea
		return
	}

	//beego.Info("=====================================切割没运行到这====================================")

	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&area)
	if err != nil || num == 0 {
		resp["errno"] = "4000"
		resp["errmsg"] = "获取数据失败"
		return
	}

	resp["data"] = &area

	json_data, err := json.Marshal(resp)
	if err != nil {
		resp["errno"] = 40002
		resp["errmsg"] = "change json data err"

		return
	}
	if err := redisData.Put("AreaJson", json_data, time.Second*3600); err != nil {
		models.MakeLogs("set area redis fail:", err)
		return
	}
	//c.RetData(resp)
	//beego.Info(json_data)

}
