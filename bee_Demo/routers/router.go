package routers

import (
	"bee_Demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//获取管理员信息
	//beego.RenderForm("admin/info", &controllers)
	//管理员账号退出
	//获取管理员列表信息
}
