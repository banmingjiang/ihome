package models

import (
	"github.com/astaxie/beego"
)

func MakeLogs(str string, msg interface{}) {
	beego.Info(str, msg)
}
