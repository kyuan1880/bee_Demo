package main

import (
	_ "bee_Demo/models"
	_ "bee_Demo/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" //导入驱动包
)

func main() {
	//模型自动生成结构
	//orm.RunSyncdb("default", false, true)
	beego.Run()
}
