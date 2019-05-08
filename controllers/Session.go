package controllers

import (
	"ihome/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type SessionController struct {
	beego.Controller
}

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

func (a *SessionController) RetData(resp map[string]interface{}) {
	a.Data["json"] = resp
	a.ServeJSON()
}

func (c *SessionController) GetSessionData() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	user := models.User{}
	resp["errno"] = models.RECODE_NODATA
	resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
	real_name := c.GetSession("real_name")
	if real_name != nil {
		user.Real_name = real_name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}

}
