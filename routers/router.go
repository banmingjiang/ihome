package routers

import (
	"ihome/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/area", &controllers.AreaController{}, "get:GetArea")
	beego.Router("/api/house", &controllers.HouseController{}, "get:GetHouseindex")

	beego.Router("/api/session", &controllers.SessionController{}, "get:GetSessionData")
	//注册页面，post方法为注册
	beego.Router("/api/reg", &controllers.RegisterController{}, "get:Register;post:Reg")
	//渲染登录页面,post做账号密码效验,delete退出登录
	beego.Router("/api/login", &controllers.LoginController{}, "get:Login;post:ToLog;delete:LogOut")
	//
	//beego.Router("/api/logout", &controllers.LogOutController{}, "post:LogOut")
}
